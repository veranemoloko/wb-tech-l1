package main

import (
	"fmt"
	"sync"
)

func sq(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := n * n
	fmt.Println(n, "*", n, "=", square)
	ch <- square
}

func sumSq(nums []int) {
	ch := make(chan int, len(nums))
	var wg sync.WaitGroup

	for _, n := range nums {
		wg.Add(1)
		go sq(n, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	sum := 0
	for sq := range ch {
		sum += sq
	}

	fmt.Println("sum of squares:", sum)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumSq(nums)
}
