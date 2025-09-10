package main

import (
	"fmt"
	"unicode"
)

func IsUniqueRunes(s string) bool {
	counter := make(map[rune]bool)

	for _, c := range s {
		lowC := unicode.ToLower(c)
		if counter[lowC] {
			return false
		}
	}
	return true
}

func main() {

	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "влИrtJh06"

	fmt.Printf("%s: %t\n", s1, IsUniqueRunes(s1))
	fmt.Printf("%s: %t\n", s2, IsUniqueRunes(s2))
	fmt.Printf("%s: %t\n", s3, IsUniqueRunes(s3))
	fmt.Printf("%s: %t\n", s4, IsUniqueRunes(s4))
}
