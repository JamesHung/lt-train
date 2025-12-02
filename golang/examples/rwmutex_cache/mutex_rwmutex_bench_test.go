package main

import (
	"sync"
	"testing"
)

// CounterMutex uses sync.Mutex.
type CounterMutex struct {
	mu sync.Mutex
	n  int64
}

func (c *CounterMutex) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *CounterMutex) Get() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// CounterRWMutex uses sync.RWMutex.
type CounterRWMutex struct {
	mu sync.RWMutex
	n  int64
}

func (c *CounterRWMutex) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

func (c *CounterRWMutex) Get() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.n
}

// Benchmark read-heavy workload: e.g., 10% writes, 90% reads.
func BenchmarkMutex_ReadHeavy(b *testing.B) {
	const readers = 8
	const writers = 2

	var wg sync.WaitGroup
	c := &CounterMutex{}

	b.ResetTimer()

	for i := 0; i < writers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				c.Inc()
			}
		}()
	}

	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				_ = c.Get()
			}
		}()
	}

	wg.Wait()
}

func BenchmarkRWMutex_ReadHeavy(b *testing.B) {
	const readers = 8
	const writers = 2

	var wg sync.WaitGroup
	c := &CounterRWMutex{}

	b.ResetTimer()

	for i := 0; i < writers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				c.Inc()
			}
		}()
	}

	for i := 0; i < readers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				_ = c.Get()
			}
		}()
	}

	wg.Wait()
}

// Benchmark write-heavy workload: e.g., counters/log-like workloads.
func BenchmarkMutex_WriteHeavy(b *testing.B) {
	var wg sync.WaitGroup
	c := &CounterMutex{}

	b.ResetTimer()

	const workers = 8
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				c.Inc()
			}
		}()
	}

	wg.Wait()
}

func BenchmarkRWMutex_WriteHeavy(b *testing.B) {
	var wg sync.WaitGroup
	c := &CounterRWMutex{}

	b.ResetTimer()

	const workers = 8
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				c.Inc()
			}
		}()
	}

	wg.Wait()
}
