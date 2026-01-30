package main

import "fmt"

func main() {

	fmt.Println(operation(5, 4, add))
	fmt.Println(operation(5, 4, subtract))
}

func add(a, b int) int {
	return a + b
}
func subtract(a, b int) int {
	return a - b
}

func operation(a, b int, f func(int, int) int) int {
	return f(a, b)
}
