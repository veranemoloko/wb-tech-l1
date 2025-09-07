package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	words := strings.Fields(s)
	// swap the i-th word with its counterpart from the end
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1] = words[len(words)-1], words[i]
	}

	return strings.Join(words, " ")

}

func main() {
	str := "dog with me on the Moon"
	fmt.Println(reverseString(str))
}
