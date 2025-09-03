package main

import "fmt"

func main() {

	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)

	nums := []int{1, 2, 3, 4, 5, 6}

	go func() {
		for _, n := range nums {
			ch1 <- n
		}
		close(ch1)
	}()

	go func() {
		for n := range ch1 {
			ch2 <- 2 * n
		}
		close(ch2)
	}()

	for n := range ch2 {
		fmt.Println(n)
	}
}
