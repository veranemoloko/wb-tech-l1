package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// workerPrint reads values from the channel and prints them
// worker prints values from the channel
func workerPrint(data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()       // notify waitgroup when worker finishes
	for d := range data { // read from channel until it's closed
		fmt.Println("- - - worker:", d)
	}
}

func main() {
	var workSeconds int
	fmt.Print("work seconds: ")
	fmt.Scan(&workSeconds)

	start := time.Now()

	// create a channel for sending data
	dataChan := make(chan int)
	var wg sync.WaitGroup

	// launch a worker
	wg.Add(1)
	go workerPrint(dataChan, &wg)

	// create a context with timeout to stop the generator after workSeconds
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(workSeconds)*time.Second)
	defer cancel()

	// data generator
	go func() {
		data := 0
		for {
			select {
			case <-ctx.Done(): // stop generator on timeout
				fmt.Println("- - timeout")
				close(dataChan) // closing channel signals worker to finish
				return
			case dataChan <- data: // send data to channel
				fmt.Println("- - pushed data", data)
				time.Sleep(500 * time.Millisecond) // simulate some delay
				data++
			}
		}
	}()

	wg.Wait() // wait for the worker to finish

	duration := time.Since(start).Seconds()
	fmt.Printf("- - program worked for %.2f seconds\n", duration)
}
