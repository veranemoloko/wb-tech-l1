package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Advantages of this solution:
1. Responsive: Ctrl+C triggers a controlled shutdown using context cancellation.
2. No race conditions: communication is handled via channels.
3. WaitGroup ensures that the main goroutine waits for all workers to finish.
*/

// it uses WG to signal completion
func workerPrint(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()       // Notify WaitGroup when worker finishes
	for d := range data { // Read from channel until it's closed
		fmt.Println("- - - Worker", id, ": data", d)
	}
}

func main() {

	var wCnt int
	fmt.Print("Count workers: ")
	fmt.Scan(&wCnt)

	// chan for sending data to workers
	wChan := make(chan int)
	var wg sync.WaitGroup

	// launch workers
	for wId := 1; wId <= wCnt; wId++ {
		wg.Add(1)
		go workerPrint(wId, wChan, &wg)
	}

	// ctx to control the data generator
	ctx, cancel := context.WithCancel(context.Background())

	// data generator - continuously sends integers to the channel
	go func() {
		data := 0
		for {
			select {
			case <-ctx.Done(): // stop generator when context is canceled
				fmt.Println("- - close Channel")
				close(wChan) // closing channel signals workers to finish
				return
			case wChan <- data: // send data to channel
				fmt.Println("- - pushed data", data)
				time.Sleep(500 * time.Millisecond) // Simulate some delay
				data++
			}
		}
	}()

	// Ctrl+с to gracefully stop the program
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT) // catch Ctrl+с
	fmt.Println("- wait for 'CTRL + C'")
	<-sigCh
	fmt.Println("- start CANCELLING workers")
	cancel() // сancel the generator context

	// wait for all workers to finish
	wg.Wait()
	fmt.Println("All workers stopped gracefully")
}
