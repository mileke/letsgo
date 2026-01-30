package main

import (
	"fmt"
)

func main() {
	x := 40
	y := 5
	fmt.Println("The value of x and y are:", x, "and", y)

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("This is the default case")
	}
}
