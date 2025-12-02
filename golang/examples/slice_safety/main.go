package main

import (
	"fmt"
	"sync"
)

func demoWithMutex() {
	fmt.Println("== sync.Mutex ==")

	s := []int{1, 2, 3}
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		mu.Lock()
		s[0] = 100
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		fmt.Println("read with lock:", s)
		mu.Unlock()
	}()

	wg.Wait()
}

func demoWithCopy() {
	fmt.Println("== copy before sharing ==")

	src := []int{1, 2, 3}
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		local := append([]int(nil), src...) // private copy for this goroutine
		local[0] = 200
		fmt.Println("writer copy:", local)
	}()

	go func() {
		defer wg.Done()
		reader := append([]int(nil), src...) // reader also gets its own copy
		fmt.Println("reader copy:", reader)
	}()

	wg.Wait()
}

func demoWithChannels() {
	fmt.Println("== channels to serialize access ==")

	s := []int{1, 2, 3}

	type op struct {
		kind string
		idx  int
		val  int
		resp chan []int
	}

	ops := make(chan op)
	var wg sync.WaitGroup
	wg.Add(1)

	// Single goroutine owns the slice; all access goes through the channel.
	go func() {
		defer wg.Done()
		for req := range ops {
			switch req.kind {
			case "set":
				s[req.idx] = req.val
			case "snapshot":
				// send a copy to avoid sharing backing array
				req.resp <- append([]int(nil), s...)
			}
		}
	}()

	ops <- op{kind: "set", idx: 0, val: 500}

	resp := make(chan []int, 1)
	ops <- op{kind: "snapshot", resp: resp}
	fmt.Println("snapshot via channel:", <-resp)

	close(ops)
	wg.Wait()
}

func demoImmutable() {
	fmt.Println("== immutable style ==")

	base := []int{1, 2, 3}

	// Each "mutation" returns a new slice, leaving the old one untouched.
	update := func(s []int, idx, val int) []int {
		next := append([]int(nil), s...)
		next[idx] = val
		return next
	}

	next := update(base, 0, 999)
	fmt.Println("base remains:", base)
	fmt.Println("new version:", next)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine sees base:", base)
	}()
	go func() {
		defer wg.Done()
		fmt.Println("goroutine sees next:", next)
	}()
	wg.Wait()
}

func main() {
	demoWithMutex()
	demoWithCopy()
	demoWithChannels()
	demoImmutable()
}
