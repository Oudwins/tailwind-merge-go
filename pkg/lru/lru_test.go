package lru

import "testing"

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
