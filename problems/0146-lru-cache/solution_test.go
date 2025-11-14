package lrucache

import "testing"

func TestLRUCacheExample(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	assertEqual(t, cache.Get(1), 1)
	cache.Put(3, 3) // evicts key 2
	assertEqual(t, cache.Get(2), -1)
	cache.Put(4, 4) // evicts key 1
	assertEqual(t, cache.Get(1), -1)
	assertEqual(t, cache.Get(3), 3)
	assertEqual(t, cache.Get(4), 4)
}

func TestLRUCacheUpdateAndCapacity(t *testing.T) {
	cache := Constructor(1)

	cache.Put(2, 1)
	assertEqual(t, cache.Get(2), 1)
	cache.Put(2, 2)
	assertEqual(t, cache.Get(2), 2)
	cache.Put(3, 3) // evict key 2
	assertEqual(t, cache.Get(2), -1)
	assertEqual(t, cache.Get(3), 3)
}

func TestLRUCacheManyOps(t *testing.T) {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4) // evict 1
	assertEqual(t, cache.Get(4), 4)
	assertEqual(t, cache.Get(3), 3)
	assertEqual(t, cache.Get(2), 2)
	assertEqual(t, cache.Get(1), -1)

	cache.Put(2, 22) // update moves 2 to front
	cache.Put(5, 5)  // evict 4
	assertEqual(t, cache.Get(2), 22)
	assertEqual(t, cache.Get(3), 3)
	assertEqual(t, cache.Get(4), -1)
	assertEqual(t, cache.Get(5), 5)
}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got=%d want=%d", got, want)
	}
}
