package main

import (
	"fmt"
)

func main() {
	x := 0

	for {
		if x >= 6 {
			break
		}
		x++
		fmt.Println(x)
	}
}
