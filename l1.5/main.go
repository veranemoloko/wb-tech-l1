package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerPrint(data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range data {
		fmt.Println("- - - Worker:", d)
	}
}

func main() {
	var seconds int
	fmt.Print("Work seconds: ")
	fmt.Scan(&seconds)

	start := time.Now()

	wChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go workerPrint(wChan, &wg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel()

	go func() {
		data := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("- - timeout")
				close(wChan)
				return
			case wChan <- data:
				fmt.Println("- - pushed data", data)
				time.Sleep(500 * time.Millisecond)
				data++
			}
		}
	}()

	wg.Wait()

	duration := time.Since(start).Seconds()
	fmt.Printf("- - program worked for %.2f seconds\n", duration)
}
