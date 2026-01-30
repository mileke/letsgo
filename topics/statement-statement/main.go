package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 40
	y := 5
	fmt.Println("The value of x and y are:", x, y)

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER than or EQUAL to x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}
