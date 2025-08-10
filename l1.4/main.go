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

func workerPrint(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range data {
		fmt.Println("- - - Worker", id, ":", d)
	}
}

func main() {
	var wCnt int
	fmt.Print("Count workers: ")
	fmt.Scan(&wCnt)

	wChan := make(chan int)
	var wg sync.WaitGroup

	for wId := 1; wId <= wCnt; wId++ {
		wg.Add(1)
		go workerPrint(wId, wChan, &wg)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		data := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("- - close Channel")
				close(wChan)
				return
			case wChan <- data:
				fmt.Println("- - pushed data", data)
				time.Sleep(500 * time.Millisecond)
				data++
			}
		}
	}()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT)
		fmt.Println("- wait for 'CTRL + C'")
		<-sigCh
		fmt.Println("- start CANCELLING workers")
		cancel()
	}()

	wg.Wait()
}
