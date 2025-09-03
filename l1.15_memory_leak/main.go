package main

import "fmt"

// 1. Memory leak: `justString = v[:100]` keeps a reference to the whole large string,
//    even though only a small substring is needed.
// 2. Global variable `justString`:
//    - lives until the program ends,
//    - potential data races in concurrent code

func createHugeString(size int) string {
	return string(make([]rune, size))
}

// avoids global variable and memory leak
func someFunc() string {
	v := createHugeString(1 << 10)
	return string([]rune(v[:100]))
}

func main() {
	result := someFunc()
	fmt.Println(result)
}
