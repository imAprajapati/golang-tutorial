package main

import "fmt"

var names = []string{"John", "Smmit", "Rhul", "Nothing"}

func array() {
	for _, name := range names {
		if name == "Nothing" {
			fmt.Println("Name Nothing in the array")
		}
	}
}
