package main

import (
	"fmt"
)

func main() {

	xi := []int{42, 43, 44, 45, 46}

	// for x := range xi {
	// 	fmt.Println("The index position is:", x, "and the value is:", xi[x])
	// }
	for i, v := range xi {
		fmt.Println("The index position is:", i, "and the value is:", v)
	}
}
