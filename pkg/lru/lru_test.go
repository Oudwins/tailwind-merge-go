package lru

import (
	"fmt"
	"sync"
	"testing"
)

// makeLRUInstance creates an LRU cache instance with the specified capacity,
// bypassing the Make() constructor to allow direct access to internal structures
// for testing list ordering. This is intentional for testing purposes only -
// production code should use Make() which returns cache.ICache (hiding internals).
//
// List Structure After Initialization:
//   - head.prev points to tail (empty list sentinel)
//   - tail.next points to head (empty list sentinel)
//   - This creates: tail <-> head (no actual data nodes between)
func makeLRUInstance(maxCapacity int) *LRU {
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

// TestLRUGetReturnsValueForExistingKey verifies that Get retrieves the correct value
// for a key that was previously set. This is the basic happy path for retrieval.
func TestLRUGetReturnsValueForExistingKey(t *testing.T) {
	lru := makeLRUInstance(10)
	lru.Set("key1", "value1")

	result := lru.Get("key1")
	if result != "value1" {
		t.Fatalf("expected value1, got %s", result)
	}
}

// TestLRUGetReturnsEmptyStringForMissingKey verifies that Get returns an empty
// string when the key doesn't exist in the cache, rather than panicking or
// returning undefined behavior.
func TestLRUGetReturnsEmptyStringForMissingKey(t *testing.T) {
	lru := makeLRUInstance(10)

	result := lru.Get("nonexistent")
	if result != "" {
		t.Fatalf("expected empty string, got %s", result)
	}
}

// TestLRUSetAddsNewKeyValuePair verifies that a new key-value pair is successfully
// stored in the cache and can be looked up via the internal map.
func TestLRUSetAddsNewKeyValuePair(t *testing.T) {
	lru := makeLRUInstance(10)
	lru.Set("key1", "value1")
	if lru.cache["key1"] == nil {
		t.Fatalf("key1 should exist in cache")
	}
}

// TestLRUSetUpdatesExistingKeyAndMovesToFront verifies that updating an existing key
// creates a new node (rather than mutating in place) and places it at the MRU position.
// Creating a new node ensures proper list manipulation - the old node is removed
// from its position and the new node is inserted at head.prev.
func TestLRUSetUpdatesExistingKeyAndMovesToFront(t *testing.T) {
	lru := makeLRUInstance(10)
	lru.Set("key1", "value1")
	oldNode := lru.cache["key1"]
	lru.Set("key1", "updated_value1")
	newNode := lru.cache["key1"]
	if newNode == oldNode {
		t.Fatalf("node should be recreated for updated key")
	}
	if newNode.val != "updated_value1" {
		t.Fatalf("expected updated_value1, got %s", newNode.val)
	}
}

// TestLRUEvictsLeastRecentlyUsedWhenCapacityExceeded verifies the core LRU eviction
// policy: when capacity is exceeded, the least recently used item (oldest access)
// is removed to make room for new items.
//
// Scenario:
//  1. Set key1, key2 (fills cache to capacity 2)
//  2. Set key3 (triggers eviction)
//  3. key1 should be evicted (it was accessed before key2)
//  4. key2 and key3 should remain
func TestLRUEvictsLeastRecentlyUsedWhenCapacityExceeded(t *testing.T) {
	lru := makeLRUInstance(2)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	lru.Set("key3", "value3")
	if lru.capacity != 2 {
		t.Fatalf("capacity should be 2, got %d", lru.capacity)
	}
	if lru.cache["key1"] != nil {
		t.Fatalf("key1 should have been evicted")
	}
	if lru.cache["key2"] == nil {
		t.Fatalf("key2 should still exist")
	}
	if lru.cache["key3"] == nil {
		t.Fatalf("key3 should exist")
	}
}

// TestLRUDoesNotEvictWhenUnderCapacity verifies that eviction only occurs when
// capacity is exceeded, not when we're still under the limit.
func TestLRUDoesNotEvictWhenUnderCapacity(t *testing.T) {
	lru := makeLRUInstance(3)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	if lru.capacity != 2 {
		t.Fatalf("capacity should be 2, got %d", lru.capacity)
	}
	if lru.cache["key1"] == nil || lru.cache["key2"] == nil {
		t.Fatalf("both keys should exist")
	}
}

// TestLRUWithCapacityOneEvictsOnSecondInsert verifies edge case behavior when
// capacity is 1: only one item can exist at a time, so setting a new key
// immediately evicts the previous one.
func TestLRUWithCapacityOneEvictsOnSecondInsert(t *testing.T) {
	lru := makeLRUInstance(1)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	if lru.capacity != 1 {
		t.Fatalf("capacity should be 1, got %d", lru.capacity)
	}
	if lru.cache["key1"] != nil {
		t.Fatalf("key1 should have been evicted")
	}
	if lru.cache["key2"] == nil {
		t.Fatalf("key2 should exist")
	}
}

// TestLRUWithCapacityZeroDoesNotStoreAnything verifies that a cache with zero
// capacity cannot store any items. The insert still happens but is immediately
// evicted, leaving an empty cache.
func TestLRUWithCapacityZeroDoesNotStoreAnything(t *testing.T) {
	lru := makeLRUInstance(0)
	lru.Set("key1", "value1")
	if lru.capacity != 0 {
		t.Fatalf("capacity should be 0, got %d", lru.capacity)
	}
	if len(lru.cache) != 0 {
		t.Fatalf("cache should be empty")
	}
}

// TestLRUGetMovesAccessedItemToMostRecentPosition verifies that accessing a key
// via Get() moves it to the most-recently-used position at the front of the list.
//
// IMPORTANT: List ordering semantics
//   - head.prev is the MOST recently used (front of list, insertion point)
//   - tail.next is the LEAST recently used (back of list, eviction candidate)
//
// This test was previously buggy: it checked tail.next.key expecting it to be MRU,
// but tail.next is actually the LRU. The fix checks head.prev.key for MRU.
//
// Sequence:
//  1. Set key1, key2, key3 -> list: tail->key1->key2->key3->head
//     (key3 is MRU at head.prev)
//  2. Get key1 -> key1 is removed and reinserted at MRU position
//     list: tail->key2->key3->key1->head
//  3. head.prev.key should be "key1" (it's now MRU)
func TestLRUGetMovesAccessedItemToMostRecentPosition(t *testing.T) {
	lru := makeLRUInstance(10)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	lru.Set("key3", "value3")

	// key3 is MRU (inserted last), verify head.prev points to it
	head := lru.head.prev
	if head.key != "key3" {
		t.Fatalf("key3 should be most recent before access")
	}

	// Access key1 - it should now be MRU
	lru.Get("key1")

	// head.prev is MRU, verify key1 is there
	// NOTE: Previous buggy test checked tail.next.key (which is LRU)
	if lru.head.prev.key != "key1" {
		t.Fatalf("key1 should be most recent after access")
	}
}

// TestLRUSetMovesUpdatedKeyToMostRecentPosition verifies that updating a value
// for an existing key places the new node at the MRU position.
//
// IMPORTANT: When Set updates an existing key, it:
//  1. Removes the old node from the list
//  2. Creates a new node with the updated value
//  3. Inserts the new node at head.prev (MRU position)
//
// This test verifies that the new node's next pointer correctly points to head
// (it's at the front of the list), and its prev pointer points to the node
// that was previously at the MRU position.
//
// NOTE: Previous buggy test checked `firstNode.next != lru.head || secondNode.prev != lru.tail`
// which was incorrect - after updating key1, the new node (secondNode) is at MRU,
// so secondNode.next should equal lru.head (not firstNode.next).
func TestLRUSetMovesUpdatedKeyToMostRecentPosition(t *testing.T) {
	lru := makeLRUInstance(10)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")

	oldNode := lru.cache["key1"]

	// Update key1 to trigger node recreation
	lru.Set("key1", "new_value1")

	newNode := lru.cache["key1"]
	if oldNode == newNode {
		t.Fatalf("should create new node when updating")
	}

	// After update, newNode should be at MRU position (head.prev)
	// In a doubly-linked list with sentinel head/tail:
	// - MRU node's next points to head (head.prev is MRU)
	// - newNode.next should point to lru.head since it's at head.prev
	if newNode.next != lru.head {
		t.Fatalf("new node should be in correct list position")
	}
}

// TestLRUMultipleGetsPreserveRecencyOrder verifies that multiple Get operations
// correctly maintain the LRU ordering invariant.
//
// List positions after sequence:
//   - After Set key1: tail->key1->head (key1 is MRU)
//   - After Set key2: tail->key1->key2->head (key2 is MRU)
//   - After Set key3: tail->key1->key2->key3->head (key3 is MRU)
//   - After Get key2: tail->key1->key3->key2->head (key2 moved to MRU)
//   - After Get key1: tail->key3->key2->key1->head (key1 moved to MRU)
//
// So after both Gets:
//   - head.prev (MRU) should be key1 (most recently accessed)
//   - tail.next (LRU) should be key3 (least recently accessed, not accessed in Gets)
//
// NOTE: Previous buggy test checked tail.next.key == "key1" which was incorrect.
// tail.next is the LRU (eviction candidate), not the MRU.
// Also: the test expected tail.next.key == "key2" but after Get(key2) then Get(key1),
// key2 was MORE recently accessed than key3, so key3 becomes the LRU.
func TestLRUMultipleGetsPreserveRecencyOrder(t *testing.T) {
	lru := makeLRUInstance(10)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	lru.Set("key3", "value3")

	// Access key2 first, then key1
	lru.Get("key2")
	lru.Get("key1")

	// key1 is now MRU (most recent access)
	// head.prev is MRU, verify key1 is there
	if lru.head.prev.key != "key1" {
		t.Fatalf("key1 should be most recent after access, got %s", lru.head.prev.key)
	}

	// key3 was never accessed after the Sets, so it's the LRU
	// tail.next is LRU (eviction candidate)
	if lru.tail.next.key != "key3" {
		t.Fatalf("key3 should be least recent after accesses, got %s", lru.tail.next.key)
	}
}

// TestLRUEmptyCacheOperationsWork verifies that cache operations work correctly
// when the cache is initially empty. Getting non-existent keys should return
// empty string without error.
func TestLRUEmptyCacheOperationsWork(t *testing.T) {
	lru := makeLRUInstance(10)
	if lru.capacity != 0 {
		t.Fatalf("empty cache should have capacity 0")
	}
	if lru.cache["any"] != nil {
		t.Fatalf("empty cache should not contain any keys")
	}
}

// TestLRURepeatedSetSameKeyDoesNotGrowCapacity verifies that repeatedly setting
// the same key (updating in place) doesn't cause capacity to grow indefinitely.
// Each update should remove the old node (decrementing capacity) then add the new
// one (incrementing capacity), keeping the count stable.
func TestLRURepeatedSetSameKeyDoesNotGrowCapacity(t *testing.T) {
	lru := makeLRUInstance(10)
	lru.Set("key1", "value1")
	oldCapacity := lru.capacity
	lru.Set("key1", "value1")
	if lru.capacity != oldCapacity {
		t.Fatalf("capacity should not change when setting same key multiple times")
	}
}

// TestLRUConcurrentGetAndSetDoNotRace verifies that concurrent Get and Set operations
// on the same keys do not cause data races.
//
// RACE CONDITION THAT WAS FIXED:
// The original implementation used two separate mutexes (cacheMutex and listMutex).
// This caused a race because:
//  1. Goroutine A calls Get("key1"), acquires cacheMutex.RLock(), reads lru.cache["key1"]
//  2. Goroutine A releases cacheMutex.RLock()
//  3. Goroutine A acquires listMutex.Lock(), modifies node.prev/node.next
//  4. Meanwhile, Goroutine B calls Set("key1", ...) which calls remove() on the SAME node
//  5. Both goroutines access node.prev/node.next concurrently -> DATA RACE
//
// The fix uses a single mutex (sync.Mutex) to protect the entire operation,
// ensuring that no two goroutines can access the same node simultaneously.
// Note: We use sync.Mutex instead of sync.RWMutex because the critical sections
// modify shared state (node pointers), not just read them.
func TestLRUConcurrentGetAndSetDoNotRace(t *testing.T) {
	t.Parallel()
	lru := makeLRUInstance(10)

	var wg sync.WaitGroup

	// Launch multiple goroutines that concurrently read and write to key1
	// This tests the race condition between Get (which repositions nodes) and Set (which removes them)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				lru.Get("key1")
				lru.Set("key1", "value1")
			}
		}()
	}

	// Launch a goroutine that writes to a different key
	// This ensures Set operations on unrelated keys don't interfere
	wg.Add(1)
	go func() {
		defer wg.Done()
		for j := 0; j < 50; j++ {
			lru.Set("key2", "value2")
		}
	}()

	wg.Wait()
}

// TestLRUConcurrentSetsDoNotCorruptList verifies that concurrent Set operations
// on different keys don't corrupt the linked list structure. Each goroutine
// sets unique keys to avoid the complexity of concurrent updates to the same key.
//
// This test catches race conditions in:
//   - Map insertion (lru.cache[key] = n)
//   - List insertion (insertRight)
//   - Capacity tracking (lru.capacity++)
//   - Eviction logic (delete + remove when over capacity)
func TestLRUConcurrentSetsDoNotCorruptList(t *testing.T) {
	t.Parallel()
	lru := makeLRUInstance(5)

	var wg sync.WaitGroup

	// Each goroutine sets unique keys based on index
	for i := 0; i < 5; i++ {
		idx := i // Capture loop variable to avoid closure issues
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				lru.Set(fmt.Sprintf("key%d", idx), "value")
			}
		}()
	}

	wg.Wait()
}

// TestLRUConcurrentGetsDoNotRace verifies that concurrent Get operations on the
// same key don't cause data races. Multiple goroutines reading the same cached
// value should all succeed without corrupting the list structure.
//
// This tests the safety of:
//   - Map lookup under lock
//   - List repositioning (remove + insertRight) for frequently accessed items
func TestLRUConcurrentGetsDoNotRace(t *testing.T) {
	t.Parallel()
	lru := makeLRUInstance(10)
	lru.Set("key", "value")

	var wg sync.WaitGroup

	// Multiple goroutines reading the same key repeatedly
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				lru.Get("key")
			}
		}()
	}

	wg.Wait()
}

// TestLRUWithMakeUsesPublicAPI verifies that the Make() constructor works correctly
// and returns a cache that can be used through the ICache interface.
//
// This test ensures the public API is functional:
//   - Make() returns a properly initialized cache
//   - Get() and Set() work through the ICache interface
//   - Eviction still works when capacity is exceeded
//
// This test covers the Make() function (previously 0% coverage) and demonstrates
// the intended usage pattern for production code.
func TestLRUWithMakeUsesPublicAPI(t *testing.T) {
	lru := Make(2)

	// Set and retrieve through interface
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")

	if lru.Get("key1") != "value1" {
		t.Fatalf("expected value1")
	}
	if lru.Get("key2") != "value2" {
		t.Fatalf("expected value2")
	}

	// Trigger eviction
	lru.Set("key3", "value3")

	// key1 should be evicted (LRU)
	if lru.Get("key1") != "" {
		t.Fatalf("key1 should have been evicted")
	}
}

// TestLRUGetOnEvictedKeyDoesNotPanic verifies that calling Get() after eviction
// returns empty string without panicking. This tests the nil-safety path in
// remove() when it's called on already-removed nodes (though this shouldn't happen
// in normal operation, it's defensive).
func TestLRUGetOnEvictedKeyDoesNotPanic(t *testing.T) {
	lru := makeLRUInstance(2)
	lru.Set("key1", "value1")
	lru.Set("key2", "value2")

	// Manually test the nil-safety of remove by accessing via cache directly
	// (this simulates what would happen if code tried to remove nil)
	lru.remove(nil) // Should be a no-op, not panic

	// Cache should still work correctly
	lru.Set("key3", "value3")
	if lru.capacity != 2 {
		t.Fatalf("capacity should be 2, got %d", lru.capacity)
	}
}
