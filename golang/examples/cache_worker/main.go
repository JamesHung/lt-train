package main

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	value     string
	expiresAt time.Time
}

// Cache has short read locks and background refresh to avoid starvation.
type Cache struct {
	mu        sync.RWMutex
	store     map[string]cacheEntry
	refreshCh chan string
	loadFn    func(string) (string, error)
	ttl       time.Duration
	wg        sync.WaitGroup
}

func NewCache(loadFn func(string) (string, error), ttl time.Duration) *Cache {
	c := &Cache{
		store:     make(map[string]cacheEntry),
		refreshCh: make(chan string, 8),
		loadFn:    loadFn,
		ttl:       ttl,
	}
	c.wg.Add(1)
	go c.worker()
	return c
}

func (c *Cache) Close() {
	close(c.refreshCh)
	c.wg.Wait()
}

// Get takes a quick read lock; if missing/expired it triggers async refresh.
func (c *Cache) Get(key string) (string, bool) {
	now := time.Now()
	c.mu.RLock()
	entry, ok := c.store[key]
	if ok && now.Before(entry.expiresAt) {
		val := entry.value
		c.mu.RUnlock()
		return val, true
	}
	c.mu.RUnlock()

	// Trigger refresh without blocking readers forever.
	select {
	case c.refreshCh <- key:
	default:
		// drop if channel is full to avoid backpressure; next Get will retry
	}

	return "", false
}

func (c *Cache) worker() {
	defer c.wg.Done()
	for key := range c.refreshCh {
		val, err := c.loadFn(key) // no locks held while fetching
		if err != nil {
			continue
		}
		c.mu.Lock()
		c.store[key] = cacheEntry{
			value:     val,
			expiresAt: time.Now().Add(c.ttl),
		}
		c.mu.Unlock()
	}
}

func main() {
	loads := make(map[string]int)
	loadFn := func(key string) (string, error) {
		loads[key]++
		time.Sleep(20 * time.Millisecond) // simulate DB/API
		return fmt.Sprintf("%s-val-%d", key, loads[key]), nil
	}

	cache := NewCache(loadFn, 120*time.Millisecond)
	defer cache.Close()

	for i := 0; i < 6; i++ {
		if v, ok := cache.Get("a"); ok {
			fmt.Println("cache hit:", v)
		} else {
			fmt.Println("cache miss; refresh scheduled")
		}
		time.Sleep(60 * time.Millisecond)
	}
}
