package main

import (
	"fmt"
)

func main() {
	//fmt.Println(true && true)
	fmt.Println(true && false)
	//fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	switch {
	case false:
		fmt.Println("This doesn't print")
	case true:
		fmt.Println("This prints")
	case true:
		fmt.Println("This doesn't print")

	}
}
