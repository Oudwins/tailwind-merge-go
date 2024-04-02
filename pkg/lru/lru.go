package lru

import (
	"sync"

	"github.com/Oudwins/tailwind-merge-go/pkg/cache"
)

type node struct {
	key  string
	val  string
	prev *node
	next *node
}

type LRU struct {
	maxCapacity int
	capacity    int
	cache       map[string]*node
	head        *node
	tail        *node
	cacheMutex  sync.RWMutex
	listMutex   sync.Mutex
}

func (lru *LRU) Get(key string) string {
	lru.cacheMutex.RLock()
	n := lru.cache[key]
	if n == nil {
		lru.cacheMutex.RUnlock()
		return ""
	}
	lru.cacheMutex.RUnlock()

	lru.listMutex.Lock()
	lru.remove(n)
	lru.insertRight(n)
	lru.listMutex.Unlock()

	return n.val
}

func (lru *LRU) Set(key, value string) {
	lru.cacheMutex.Lock()
	if n := lru.cache[key]; n != nil {
		lru.remove(n)
	}
	n := &node{key: key, val: value}
	lru.cache[key] = n
	lru.cacheMutex.Unlock()
	lru.listMutex.Lock()
	lru.insertRight(n)
	lru.listMutex.Unlock()
	// evict

	lru.listMutex.Lock()
	if lru.capacity > lru.maxCapacity {
		delete(lru.cache, lru.tail.next.key)
		lru.remove(lru.tail.next)
	}
	lru.listMutex.Unlock()

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
