package lru

import (
	"strconv"
	"sync"
	"testing"
)

func TestLRUEvictsLeastRecentlyUsedWhenCapacityExceeded(t *testing.T) {
	cache := Make(2)

	cache.Set("first", "1")
	cache.Set("second", "2")
	cache.Set("third", "3")

	if got := cache.Get("first"); got != "" {
		t.Fatalf("expected first key to be evicted, got %q", got)
	}
	if got := cache.Get("second"); got != "2" {
		t.Fatalf("expected second key to remain, got %q", got)
	}
	if got := cache.Get("third"); got != "3" {
		t.Fatalf("expected third key to remain, got %q", got)
	}
}

func TestLRUConcurrentGetSet(t *testing.T) {
	cache := Make(32)

	for i := 0; i < 32; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, key)
	}

	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				key := strconv.Itoa((id + j) % 64)
				cache.Set(key, key)
				cache.Get(key)
			}
		}(i)
	}

	wg.Wait()
}
