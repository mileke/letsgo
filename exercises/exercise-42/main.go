package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//create an array which holds 5 values of type int
	ai := [5]int{}
	//assign values to each index position
	//range over the array and print the values and the type
	for i := range ai {
		ai[i] = rand.Intn(50)
		fmt.Printf("The value of the array at position %v is: %v and the type is %T\n", i, ai[i], ai[i])
	}
}
