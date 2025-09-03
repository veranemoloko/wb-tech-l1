package main

import "fmt"

func main() {
	myFriends := map[string]struct{}{
		"Alisa":    {},
		"Lena":     {},
		"Veronica": {},
		"Artem":    {},
		"Sava":     {},
	}
	herFriends := map[string]struct{}{
		"Vera":     {},
		"Lena":     {},
		"Veronica": {},
		"Max":      {},
		"Sava":     {},
	}

	ourFriends := make(map[string]struct{})

	for name := range myFriends {
		if _, ok := herFriends[name]; ok {
			ourFriends[name] = struct{}{}
		}
	}

	fmt.Println(ourFriends)
}
