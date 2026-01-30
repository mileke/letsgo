package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for i, v := range m {
		fmt.Printf("The key %v has the value %v\n", i, v)
	}
}
