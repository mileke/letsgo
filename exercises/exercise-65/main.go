package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
}

func factorial(a int) int {
	if a <= 1 {
		return 1
	}
	return a * factorial(a-1)
}
