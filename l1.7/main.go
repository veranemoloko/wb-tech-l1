package main

import (
	"fmt"
	"sync"
)

type squareMap struct {
	mu sync.Mutex
	mp map[int]int
}

func (sm *squareMap) Write(key int, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.mp[key] = value
}

func main() {
	sm := squareMap{mp: make(map[int]int)}
	var wg sync.WaitGroup

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range nums {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Write(i, i*i)
		}(i)
	}

	wg.Wait()

	for k, v := range sm.mp {
		fmt.Printf("%d: %d\n", k, v)
	}

}
