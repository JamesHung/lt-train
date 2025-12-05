package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func demoGoroutinesAndWaitGroup() {
	fmt.Println("== goroutines + sync.WaitGroup ==")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 30 * time.Millisecond)
			fmt.Printf("worker %d done\n", id)
		}(i)
	}
	wg.Wait()
}

func demoUnbufferedChannel() {
	fmt.Println("== unbuffered channel (sync handoff) ==")

	ch := make(chan string) // cap=0
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		msg := "ping"
		fmt.Println("sending:", msg)
		ch <- msg // blocks until receiver ready
		fmt.Println("sent:", msg)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(20 * time.Millisecond)
		v := <-ch
		fmt.Println("received:", v)
	}()

	wg.Wait()
}

func demoBufferedChannel() {
	fmt.Println("== buffered channel (cap > 0) ==")

	jobs := make(chan int, 2) // buffered channel
	jobs <- 1                 // does not block (buffer has space)
	jobs <- 2                 // still does not block (buffer now full)
	fmt.Printf("buffered len=%d cap=%d\n", len(jobs), cap(jobs))

	close(jobs)
	for j := range jobs {
		fmt.Printf("drained job %d\n", j)
	}
}

func demoChannelClose() {
	fmt.Println("== channel close semantics ==")

	ch := make(chan int, 1)
	ch <- 7
	close(ch)

	v, ok := <-ch
	fmt.Printf("first read after close: v=%d ok=%v\n", v, ok) // gets queued value

	v2, ok2 := <-ch
	fmt.Printf("second read after close: v=%d ok=%v\n", v2, ok2) // zero value, ok=false
}

func demoSelectMultiplex() {
	fmt.Println("== select multiplexing ==")

	ch1 := make(chan int, 1)
	ch2 := make(chan int) // unbuffered; send will block unless chosen
	ch1 <- 42

	for {
		select {
		case v := <-ch1:
			fmt.Println("got from ch1:", v)
			go func() {
				ch2 <- 10
			}()
			// case ch2 <- 10:
			fmt.Println("sent to ch2")
		case v, _ := <-ch2:
			fmt.Println("got from ch2:", v)
			return
		default:
			fmt.Println("meet default")
			time.Sleep(time.Millisecond * 10)

		}
	}
}

func demoSelectDefault() {
	fmt.Println("== select with default (non-blocking) ==")

	ch := make(chan int)
	select {
	case v := <-ch:
		fmt.Println("got:", v)
	default:
		fmt.Println("no message")
	}
}

func demoSelectTimeout() {
	fmt.Println("== select with timeout ==")

	result := make(chan string, 1)
	go func() {
		time.Sleep(40 * time.Millisecond)
		result <- "finished work"
	}()

	select {
	case msg := <-result:
		fmt.Println("select received:", msg)
	case <-time.After(30 * time.Millisecond):
		fmt.Println("select timed out")
	}
}

func demoMutex() {
	fmt.Println("== sync.Mutex for shared state ==")

	var mu sync.Mutex
	total := 0
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			total++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("total after mutex-protected increments:", total)
}

func demoContextCancel() {
	fmt.Println("== context.WithCancel ==")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker stopped:", ctx.Err())
				return
			default:
				time.Sleep(15 * time.Millisecond)
			}
		}
	}()

	ctx.Done()

	time.Sleep(35 * time.Millisecond)
	cancel()
	<-done
}

func demoContextTimeout() {
	fmt.Println("== context.WithTimeout ==")

	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("work finished")
	case <-ctx.Done():
		fmt.Println("timeout triggered:", ctx.Err())
	}
}

func demoHonorCancellationInLoop() {
	fmt.Println("== honoring cancellation in loop ==")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(40 * time.Millisecond)
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("loop exit on cancel")
			return
		default:
			// simulate work
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func taskWorker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			fmt.Println("doing something")
		}
	}
}

func demoTaskWorkerWithCancel() {
	fmt.Println("== task worker with context cancel (ticker) ==")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go taskWorker(ctx, &wg)
	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}

func demoProducerConsumerWithTimeout() {
	fmt.Println("== producer/consumer with timeout ==")

	ch := make(chan int, 5)
	defer close(ch)
	timeOut := 150 * time.Millisecond

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {

		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	go func() {
		timer := time.NewTimer(timeOut)
		defer timer.Stop()
		defer wg.Done()

		for {
			select {
			case v := <-ch:
				fmt.Println("received: ", v)
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(timeOut)
			case <-timer.C:
				fmt.Printf("Timeouted")
				return
			}
		}

	}()

	wg.Wait()

	// ch := make(chan int, 5)
	// timeout := 150 * time.Millisecond

	// var wg sync.WaitGroup
	// wg.Add(2)

	// // producer
	// go func() {
	// 	defer wg.Done()
	// 	defer close(ch)
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 		time.Sleep(30 * time.Millisecond)
	// 	}
	// }()

	// // consumer with timeout
	// go func() {
	// 	defer wg.Done()
	// 	timer := time.NewTimer(timeout)
	// 	defer timer.Stop()

	// 	for {
	// 		select {
	// 		case v, ok := <-ch:
	// 			if !ok {
	// 				fmt.Println("channel closed, consumer exited")
	// 				return
	// 			}
	// 			fmt.Println("consumed:", v)
	// 			if !timer.Stop() {
	// 				<-timer.C
	// 			}
	// 			timer.Reset(timeout)
	// 		case <-timer.C:
	// 			fmt.Println("no more data (consumer timeout)")
	// 			return
	// 		}
	// 	}
	// }()

	// wg.Wait()
}

func main() {
	demoGoroutinesAndWaitGroup()
	demoUnbufferedChannel()
	demoBufferedChannel()
	demoChannelClose()
	demoSelectMultiplex()
	demoSelectDefault()
	demoSelectTimeout()
	demoMutex()
	demoContextCancel()
	demoContextTimeout()
	demoHonorCancellationInLoop()
	demoTaskWorkerWithCancel()
	demoProducerConsumerWithTimeout()
}
