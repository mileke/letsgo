package main

import "fmt"

func main() {
	fmt.Println(factorial(4))

	fmt.Println(factLoop(4))
}

func factorial(a int) int {
	if a <= 1 {
		return 1
	}
	return a * factorial(a-1)
}

func factLoop(a int) int {
	total := 1
	for a > 0 {
		fmt.Println(a)
		total *= a
		a--
	}
	return total
}
