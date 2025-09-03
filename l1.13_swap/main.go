package main

import "fmt"

func main() {

	x, y := -2, 22

	fmt.Println("x=", x, "y=", y)

	x = x + y
	y = x - y
	x = x - y

	fmt.Println("focus-pocus...")
	fmt.Println("x=", x, "y=", y)
	fmt.Println("pocus-focus...")

	x = x ^ y
	y = x ^ y
	x = x ^ y

	fmt.Println("x=", x, "y=", y)

	x, y = y, x

	fmt.Println("and...")
	fmt.Println("x=", x, "y=", y)
}
