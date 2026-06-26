// Package lru implements a thread-safe Least Recently Used (LRU) cache with O(1)
// get, set, and eviction operations.
//
// The LRU cache maintains a fixed maximum capacity and automatically evicts the least
// recently accessed items when new items are added beyond capacity. This is useful
// for caching expensive computations where only the most recent results need to be retained.
//
// Concurrency Model:
//
// The LRU cache uses a single mutex (sync.Mutex) to protect all operations rather than
// separate read/write locks. This design choice is critical for correctness:
//
//   - The cache map stores pointers to nodes, and both Get and Set operations modify
//     the node's prev/next fields when reordering the linked list
//
//   - Using separate mutexes would allow one goroutine to read node pointers while another
//     modifies them, causing data races when:
//     1. Get() reads lru.cache[key] under cacheMutex, releases it, then modifies the node
//     2. Concurrent Set() modifies the same node's links under listMutex
//     3. Both access n.prev/n.next simultaneously (race condition!)
//
//   - A single mutex ensures atomicity of the entire "lookup and reposition" operation
//     in Get(), and atomicity of "check existing, create node, insert, maybe evict"
//     in Set()
//
// Data Structure:
//
// The cache uses a doubly-linked list with sentinel head/tail nodes:
//
//	[tail] <-> [LRU items...] <-> [MRU item] <-> [head]
//
//	- tail.next points to the least recently used item (eviction candidate)
//	- head.prev points to the most recently used item
//	- head and tail are sentinel nodes that never hold actual data
//	- New items are inserted after head (front = most recent)
//	- Eviction removes from tail.next (back = least recent)
package lru

import (
	"sync"

	"github.com/Oudwins/tailwind-merge-go/pkg/cache"
)

// node represents a single entry in the LRU cache's doubly-linked list.
// Each node stores a key-value pair and maintains prev/next pointers for
// O(1) removal and repositioning operations.
type node struct {
	key  string // The cache key for lookup
	val  string // The cached value
	prev *node  // Previous node in the linked list
	next *node  // Next node in the linked list
}

// LRU implements a concurrent-safe Least Recently Used cache with a fixed capacity.
// Items are evicted from the tail (least recently used) when capacity is exceeded.
// The linked list maintains most-recently-used items near head, least-recently-used
// near tail for efficient ordering.
type LRU struct {
	maxCapacity int              // Maximum number of items allowed in the cache
	capacity    int              // Current number of items in the cache
	cache       map[string]*node // Hash map for O(1) key lookups to node pointers
	head        *node            // Sentinel node marking the front (MRU position)
	tail        *node            // Sentinel node marking the back (LRU position)
	mutex       sync.Mutex       // Single mutex protecting all cache and list operations
}

// Get retrieves a value from the cache and marks it as most recently used.
// If the key exists, the node is moved to the front of the list (after head)
// so it becomes the most recently used item. If the key doesn't exist, an
// empty string is returned. This operation is O(1).
//
// Thread Safety: Uses mutex.Lock() to ensure the entire lookup and list
// repositioning is atomic, preventing races with concurrent Set operations.
func (lru *LRU) Get(key string) string {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	n := lru.cache[key]
	if n == nil {
		return ""
	}
	// Move node to MRU position (front of list, right after head)
	lru.remove(n)
	lru.insertRight(n)
	return n.val
}

// Set stores a key-value pair in the cache, evicting the LRU item if necessary.
// If the key already exists, the old node is removed from its current position
// before the new value is inserted at the MRU position. The capacity counter is
// incremented before checking against maxCapacity, ensuring eviction happens when
// we exceed the limit. This operation is O(1).
//
// Eviction Strategy:
//   - When capacity > maxCapacity after insertion, removes tail.next (LRU item)
//   - This maintains the invariant: capacity never exceeds maxCapacity
//
// Thread Safety: Uses mutex.Lock() to ensure the entire operation (check existing,
// create node, insert, maybe evict) is atomic.
func (lru *LRU) Set(key, value string) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	// If key exists, remove the old node before creating the new one
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}
	n := &node{key: key, val: value}
	lru.cache[key] = n
	lru.insertRight(n) // Insert at MRU position (after head)

	// Evict LRU item if we've exceeded capacity
	if lru.capacity > lru.maxCapacity {
		delete(lru.cache, lru.tail.next.key)
		lru.remove(lru.tail.next)
	}
}

// insertRight adds a node at the front of the doubly-linked list (most recently used
// position). This is called after node creation for Set operations, and after Get
// operations to reposition accessed nodes.
//
// List Structure Before:
//
//	[tail] -> ... -> [prev=head.prev] -> [head] <- [nxt=head.next] <- ... <- [tail]
//
//	List Structure After:
//
//	[tail] -> ... -> [head] <- [n] <- [head] (n becomes head.prev, the MRU)
//
// This function must be called with mutex already locked by the caller.
func (lru *LRU) insertRight(n *node) {
	prev := lru.head.prev
	prev.next = n
	n.prev = prev
	n.next = lru.head
	lru.head.prev = n
	lru.capacity++
}

// remove detaches a node from the doubly-linked list and resets its pointers.
// This is used to reposition existing nodes (in Get) or remove them entirely
// (during eviction). The node's prev/next links are set to nil to ensure
// any concurrent access to the node cannot corrupt the list structure.
//
// Note: This function accesses lru.capacity, which is why it requires mutex
// protection. A data race was occurring when Get() acquired node pointers under
// cacheMutex but then modified them under listMutex while Set() held listMutex
// and called remove() on the same nodes.
//
// This function must be called with mutex already locked by the caller.
func (lru *LRU) remove(n *node) {
	if n == nil {
		return
	}
	prev := n.prev
	nxt := n.next
	prev.next = nxt
	nxt.prev = prev
	n.prev = nil
	n.next = nil
	lru.capacity--
}

// Make creates a new LRU cache with the specified maximum capacity.
// The returned cache implements the cache.ICache interface and is safe for
// concurrent use. Items are evicted using LRU policy when capacity is exceeded.
//
// Initial State:
//   - Empty cache with 0 items
//   - head and tail sentinel nodes initialized with head.prev = tail, tail.next = head
//   - This creates an empty list: tail <-> head
func Make(maxCapacity int) cache.ICache {
	head := &node{}
	tail := &node{}
	tail.next = head
	head.prev = tail
	return &LRU{
		maxCapacity: maxCapacity,
		capacity:    0,
		cache:       make(map[string]*node),
		head:        head,
		tail:        tail,
	}
}
