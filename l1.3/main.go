package main

import "fmt"

func squareNum(num int, ch chan<- int) {
	ch <- num * num
}

func sumSquares(nums []int) int {
	ch := make(chan int)

	for _, n := range nums {
		go squareNum(n, ch)
	}

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += <-ch
	}
	return sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	sum := sumSquares(numbers)
	fmt.Println("sum of squares:", sum)
}
