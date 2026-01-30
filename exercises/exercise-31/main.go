package main

import (
	"fmt"
)

func main() {
	x := 0

	for x < 5 {
		x++
		fmt.Println(x)
	}

	for {
		if x >= 6 {
			break
		}
		x++
		fmt.Println(x)
	}
}
