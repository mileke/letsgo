package main

import "fmt"

func main() {
	fmt.Println(doOperation(5, 4, add))
	fmt.Println(doOperation(5, 4, subtract))
}

func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}
func doOperation(a int, b int, f func(c int, d int) int) int {
	return f(a, b)
}
