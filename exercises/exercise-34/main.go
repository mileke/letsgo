package main

import (
	"fmt"
)

func main() {

	for x := range 5 {
		for y := range 5 {
			fmt.Printf("x: %d, y: %d\n", x, y)
		}
	}
}
