package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Println("The value of x is:", x)

	switch {
	case x <= 100:
		fmt.Println("x is between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("x is between 101 and 200")
	case x > 200:
		fmt.Println("x is greater than 200")
	default:
		fmt.Println("This is the default case")
	}
}
