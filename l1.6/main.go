package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func stopWithFlag() {
	fmt.Println("\n- - - stop with flag")
	ch := make(chan int)

	go func() {
		fmt.Println("- go start")
		for val := range ch {
			if val == -1 {
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
	ch <- -1
	time.Sleep(300 * time.Millisecond)

}

func stopWithChannelClose() {
	fmt.Println("\n- - - stop with channel close")
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("- go start")
		for val := range ch {
			fmt.Println(val)
		}
		fmt.Println("- go finish")
	}()

	ch <- 1
	close(ch)
	wg.Wait()
}

func stopWithDoneChannel() {
	fmt.Println("\n- - - stop with channel done")
	done := make(chan struct{})

	go func() {
		fmt.Println("- go start")
		for {
			select {
			case <-done:
				fmt.Println("- go finish")
				return
			default:
				fmt.Println(2)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	done <- struct{}{}
}

func stopWithContextCancel() {
	fmt.Println("\n- - - stop with context cancel")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		fmt.Println("- go start")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("- go finish", ctx.Err())
				return
			default:
				fmt.Println(3)
				time.Sleep(50 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(100 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func stopWithContextTimeout() {
	fmt.Println("\n- - - stop with context timeout")
	ctx, cancel := context.WithTimeout(context.Background(), 222*time.Millisecond)
	defer cancel()

	go func(ctx context.Context) {
		fmt.Println("- go start")
		for {
			select {
			case <-ctx.Done():
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

func main() {
	stopWithFlag()
	stopWithDoneChannel()
	stopWithChannelClose()
	stopWithContextCancel()
	stopWithContextTimeout()
}
