package lru

import "sync"

type node struct {
	key  string
	val  string
	prev *node
	next *node
}

type Cache interface {
	Get(string) string
	Set(string, string)
}

type LRU struct {
	maxCapacity int
	capacity    int
	cache       map[string]*node
	head        *node
	tail        *node
	mutex       sync.Mutex
}

func (lru *LRU) Get(key string) string {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	n := lru.cache[key]
	if n == nil {
		return ""
	}

	lru.remove(n)
	lru.insertRight(n)

	return n.val
}

func (lru *LRU) Set(key, value string) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}

	n := &node{key: key, val: value}
	lru.cache[key] = n
	lru.insertRight(n)

	// evict
	if lru.capacity > lru.maxCapacity {
		delete(lru.cache, lru.tail.next.key)
		lru.remove(lru.tail.next)
	}

}

func (lru *LRU) insertRight(n *node) {
	prev := lru.head.prev
	prev.next = n
	n.prev = prev
	n.next = lru.head
	lru.head.prev = n
}

func (lru *LRU) remove(n *node) {
	prev := n.prev
	nxt := n.next
	prev.next = nxt
	nxt.prev = prev
	n.prev = nil
	n.next = nil
	lru.capacity--
}

func Make(maxCapacity int) Cache {
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
