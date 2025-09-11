package main

import "fmt"

// createHugeString simulates a large string
func createHugeString(size int) string {
	return string(make([]rune, size))
}

// someFunc returns only the needed substring
func someFunc() string {
	v := createHugeString(1 << 10) // large string

	// Answer: slicing v[:100] keeps reference to the whole string, causing memory retention.
	// Fix: create a new string containing only the needed part.
	sub := string([]rune(v[:100])) // copy only the first 100 runes
	return sub
}

func main() {
	result := someFunc()
	fmt.Println(result)
}
