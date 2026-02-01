package main

import "fmt"

func main() {
	fmt.Println(doOperation(square, 7))
}

/* Another way to define that function is to use types
type process func(n int) int

func doesOperation(f process, a int) int {
	x := f(a)
	return x
}*/

func doOperation(f func(n int) int, a int) int {
	x := f(a)
	return x
}

func square(n int) int {
	return n * n
}
