package main

import (
	"fmt"
	"sync"
	"time"
)

// workerPrint reads data from the channel and prints it
func workerPrint(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range data {
		fmt.Printf("Worker %d: %d\n", id, d)
	}
}

func main() {
	var wCnt int
	fmt.Print("Count workers: ")
	fmt.Scan(&wCnt)

	dataCh := make(chan int)
	var wg sync.WaitGroup

	// Start N workers
	for i := 1; i <= wCnt; i++ {
		wg.Add(1)
		go workerPrint(i, dataCh, &wg)
	}

	// Data generator - sends integers into the channel
	go func() {
		n := 0
		for {
			dataCh <- n
			n++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}
