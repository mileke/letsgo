package main

import "fmt"

// Function that takes another function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Functions to be used as arguments
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	result1 := applyOperation(5, 3, add)      // Passing 'add' function as argument
	result2 := applyOperation(8, 4, subtract) // Passing 'subtract' function as argument

	fmt.Println(result1) // Output: 8
	fmt.Println(result2) // Output: 4
}
