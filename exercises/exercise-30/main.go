package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(5)
	for i := 1; i <= 42; i++ {
		switch x {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		default:
			fmt.Println("Unknown Number")
		}
		fmt.Println("This loop has run", i, "times")
	}
}
