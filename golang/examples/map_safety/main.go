package main

import (
	"fmt"
	"sync"
)

// demoUnsafe shows the pattern that can panic with "concurrent map writes".
// It is not called from main to keep the example runnable without panic.
func demoUnsafe() {
	fmt.Println("== unsafe: concurrent access (may panic) ==")
	m := make(map[int]int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		m[1] = 10 // write
	}()
	go func() {
		defer wg.Done()
		_ = m[1] // read
	}()
	wg.Wait()
}

func demoWithMutex() {
	fmt.Println("== sync.Mutex ==")
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	written := make(chan struct{})
	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		m[1] = 10
		mu.Unlock()
		close(written)
	}()

	go func() {
		defer wg.Done()
		<-written // ensure write happened before read for demo output
		mu.Lock()
		fmt.Println("read with lock:", m[1])
		mu.Unlock()
	}()

	wg.Wait()
}

func demoWithRWMutex() {
	fmt.Println("== sync.RWMutex (read-heavy) ==")
	m := make(map[int]int)
	var mu sync.RWMutex
	var wg sync.WaitGroup
	written := make(chan struct{})

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		m[1] = 10
		mu.Unlock()
		close(written)
	}()

	// Readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			<-written
			mu.RLock()
			fmt.Printf("reader %d saw %d\n", id, m[1])
			mu.RUnlock()
		}(i)
	}

	wg.Wait()
}

func demoWithSyncMap() {
	fmt.Println("== sync.Map (many goroutines, rare writes) ==")
	var m sync.Map

	var wg sync.WaitGroup
	written := make(chan struct{})
	wg.Add(2)

	go func() {
		defer wg.Done()
		m.Store(1, 10)
		close(written)
	}()

	go func() {
		defer wg.Done()
		<-written
		if v, ok := m.Load(1); ok {
			fmt.Println("sync.Map read:", v.(int))
		} else {
			fmt.Println("sync.Map read: missing")
		}
	}()

	wg.Wait()
}

func main() {
	// demoUnsafe() // Uncomment to see the potential panic or race.
	demoWithMutex()
	demoWithRWMutex()
	demoWithSyncMap()
}
