package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeCache struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewSafeCache() *SafeCache {
	return &SafeCache{
		store: make(map[string]string),
	}
}

// Get allows concurrent readers.
func (c *SafeCache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.store[key]
}

// Set requires exclusive access.
func (c *SafeCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func main() {
	cache := NewSafeCache()

	// Writer goroutine.
	go func() {
		for i := 0; i < 5; i++ {
			v := fmt.Sprintf("val-%d", i)
			cache.Set("x", v)
			fmt.Println("writer wrote:", v)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Many readers.
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				time.Sleep(time.Duration(rand.Intn(80)+20) * time.Millisecond)
				v := cache.Get("x")
				fmt.Printf("reader %d got: %v\n", id, v)
			}
		}(i)
	}

	wg.Wait()
}
