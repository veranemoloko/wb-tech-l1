package main

import (
	"fmt"
	"sync"
)

func sq(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n, "*", n, "=", n*n)
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup

	for _, n := range nums {
		wg.Add(1)
		go sq(n, &wg)
	}

	wg.Wait()
}
