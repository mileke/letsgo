package main

import (
	"fmt"
)

func main() {
	//create an array which holds 5 values of type int
	xi := make([]int, 10)
	//assign values to each index position
	//range over the array and print the values and the type
	x := 42
	for i := range xi {
		xi[i] = (x + i)
		fmt.Printf("The value of the slice at position %v is: %v and the type is %T\n", i, xi[i], xi[i])
	}
	fmt.Println(xi)
	fmt.Println(append(xi[:3], xi[6:]...))
}
