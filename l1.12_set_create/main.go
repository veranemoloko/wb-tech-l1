package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	setStr := make(map[string]struct{})

	for _, s := range data {
		setStr[s] = struct{}{}
	}

	fmt.Println(setStr)
}
