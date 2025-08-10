package main

import (
	"fmt"
	"sync"
)

func squareNum(num int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	sq := num * num
	fmt.Printf("%d * %d = %d\n", num, num, sq)
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	var waitGroup sync.WaitGroup

	for num := range nums {
		waitGroup.Add(1)
		go squareNum(num, &waitGroup)
	}

	waitGroup.Wait()
}
