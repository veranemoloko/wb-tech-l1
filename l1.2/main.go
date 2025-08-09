package main

import (
	"fmt"
	"sync"
)

func squareNum(num int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	pow2 := num * num
	fmt.Printf("%d * %d = %d\n", num, num, pow2)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	var waitGroup sync.WaitGroup

	for _, num := range nums {
		waitGroup.Add(1)
		go squareNum(num, &waitGroup)
	}

	waitGroup.Wait()
}
