package main

import (
	"fmt"
)

func removeWithOrder(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// more faster
func removeWithoutOrder(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice:", s1)

	s1 = removeWithOrder(s1, 2)
	fmt.Println("Remove with order:", s1)

	s2 = removeWithoutOrder(s2, 2)
	fmt.Println("Remove without order:", s2)
}
