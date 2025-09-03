package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := true
	ch := make(chan int)

	catInABag := []interface{}{
		"wombat",
		22,
		ch,
		b,
	}

	// Determine the type of each element using a type switch
	for _, x := range catInABag {
		switch v := x.(type) {
		case string:
			fmt.Println(v, "- string")
		case int:
			fmt.Println(v, "- int")
		case chan int:
			fmt.Println(v, "- chan")
		case bool:
			fmt.Println(v, "- bool")
		}
	}

	// Determine the type of each element using reflection
	for _, x := range catInABag {
		t := reflect.TypeOf(x)
		v := reflect.ValueOf(x)

		switch t.Kind() {
		case reflect.String:
			fmt.Println(v.String(), "- string, len")
		case reflect.Int:
			fmt.Println(v.Int(), "- int")
		case reflect.Bool:
			fmt.Println(v.Bool(), "- bool")
		case reflect.Chan:
			fmt.Println(v, "- chan")
		}
	}
}
