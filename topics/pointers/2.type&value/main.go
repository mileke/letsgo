package main

import "fmt"

func main() {
	x := 21
	y := &x

	fmt.Printf("The value of y is: %v \n", y)
	fmt.Printf("The type of y is: %T", y)
}
