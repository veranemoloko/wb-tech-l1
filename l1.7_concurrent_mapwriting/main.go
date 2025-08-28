package main

import (
	"fmt"
	"sync"
)

func main() {
	// declare a concurrent-safe map
	var sm sync.Map
	var wg sync.WaitGroup

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// launch a goroutine for each number to write its square to sm
	for i := range nums {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Store(i, i*i) // store square in sm
		}(i) // pass i as parameter to avoid closure capture issue
	}

	wg.Wait()

	sm.Range(func(key, value any) bool {
		fmt.Printf("%d: %d\n", key, value)
		return true // return true to continue iteration
	})
}
