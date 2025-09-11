package main

import (
	"fmt"
)

func setBitInt64(num *int64, i uint, bit uint8) {
	switch bit {
	case 1:
		*num = *num | (1 << i)
	case 0:
		*num = *num &^ (1 << i)
	default:
		fmt.Println("invalid bit value")
	}
}

func main() {
	var num int64 = 259
	fmt.Printf("number %d: %14b\n", num, num)

	setBitInt64(&num, 0, 0)
	fmt.Printf("zero bit to 0: %11b\n", num)

	setBitInt64(&num, 2, 1)
	fmt.Printf("second bit to 1: %b\n", num)

}
