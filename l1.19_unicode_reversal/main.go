package main

import "fmt"

// reverseUnicode takes a UTF-8 string, reverses it, and returns the reversed string.
func reverseUnicode(s string) string {
	// convert string to rune slice (handles Unicode)
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		// swap rune at position i with its counterpart from the end
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}

func main() {
	str := "главрыба"
	reversed := reverseUnicode(str)
	fmt.Println(reversed)
}
