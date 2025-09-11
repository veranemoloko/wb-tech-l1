package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Comparison of goroutine stopping methods:

1. stop with flag
   - pros:
     * simple to implement
     * easy to understand
     * works with unbuffered channels
   - cons:
     * requires special sentinel value (-1)
     * all readers must handle the special value
     * can be error-prone if sentinel value collides with real data

2. stop with channel close
   - pros:
     * no special sentinel values needed
     * automatically signals all receivers
   - cons:
     * channel cannot be reused after close
     * only suitable for "one-time" termination

3. stop with done channel
   - pros:
     * clear separation of control and data
     * works with multiple goroutines sharing the same done signal
     * allows flexible termination without touching data channel
   - cons:
     * extra channel allocation
     * must use select to listen for done signal

4. stop with context cancel
   - pros:
     * context can carry deadlines, cancellations, and values
     * works well with multiple goroutines
   - cons:
     * slightly more boilerplate
     * context must be passed explicitly to each goroutine

5. stop with context timeout
   - pros:
     * automatic stop after a specified duration
     * reduces boilerplate in timed operations
     * integrates with context cancellation
   - cons:
     * fixed timeout may not match actual work duration
     * requires careful timeout selection

6. stop with runtime.Goexit()
   - pros:
     * immediately stops the current goroutine
     * all defer statements in the goroutine are executed
     * useful for emergency termination within the goroutine itself
   - cons:
     * only stops the current goroutine
     * no control over other goroutines
     * rarely used in idiomatic Go
     * code after Goexit() in the same goroutine will never run
*/

func main() {
	stopWithFlag()
	stopWithDoneChannel()
	stopWithChannelClose()
	stopWithContextCancel()
	stopWithContextTimeout()
	stopWithGoexit()
}

// stopWithFlag demonstrates stopping a goroutine by sending a special value (-1) through the channel
func stopWithFlag() {
	fmt.Println("\n- - - stop with flag")
	ch := make(chan int)

	go func() {
		fmt.Println("- go start")
		for val := range ch { // read values from the channel
			if val == -1 { // special value signals termination
				fmt.Println("- go finish")
				return
			}
			fmt.Println(val)
		}
	}()

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- -1 // send termination signal
	time.Sleep(300 * time.Millisecond)
}

// stopWithChannelClose demonstrates stopping a goroutine by closing the channel
func stopWithChannelClose() {
	fmt.Println("\n- - - stop with channel close")
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("- go start")
		for val := range ch { // reading from channel until it is closed
			fmt.Println(val)
		}
		fmt.Println("- go finish")
	}()

	ch <- 1
	close(ch) // closing channel signals goroutine to exit
	wg.Wait()
}

// stopWithDoneChannel demonstrates using a dedicated done channel to signal goroutine termination
func stopWithDoneChannel() {
	fmt.Println("\n- - - stop with channel done")
	done := make(chan struct{})

	go func() {
		fmt.Println("- go start")
		for {
			select {
			case <-done: // exit when done signal received
				fmt.Println("- go finish")
				return
			default:
				fmt.Println(2)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	done <- struct{}{} // send signal to stop
}

// stopWithContextCancel demonstrates stopping a goroutine using context cancellation
func stopWithContextCancel() {
	fmt.Println("\n- - - stop with context cancel")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		fmt.Println("- go start")
		for {
			select {
			case <-ctx.Done(): // exit when context is cancelled
				fmt.Println("- go finish", ctx.Err())
				return
			default:
				fmt.Println(3)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(100 * time.Millisecond)
	cancel() // cancel context to stop goroutine
	time.Sleep(100 * time.Millisecond)
}

// stopWithContextTimeout demonstrates stopping a goroutine using context timeout
func stopWithContextTimeout() {
	fmt.Println("\n- - - stop with context timeout")
	ctx, cancel := context.WithTimeout(context.Background(), 222*time.Millisecond)
	defer cancel()

	go func(ctx context.Context) {
		fmt.Println("- go start")
		for {
			select {
			case <-ctx.Done(): // exit when context times out
				fmt.Println("- go finish", ctx.Err())
				return
			default:
				fmt.Println(4)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(300 * time.Millisecond)
}

// stopWithGoexit demonstrates stopping a goroutine using runtime.Goexit
func stopWithGoexit() {
	fmt.Println("\n- - - stop with runtime.Goexit")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("- go start")
		runtime.Goexit() // immediately stops the current goroutine
		fmt.Println("- this line will never be printed")
	}()

	wg.Wait()
	fmt.Println("- goroutine stopped using Goexit")
}
