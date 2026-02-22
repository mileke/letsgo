package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNumbers interface {
	constraints.Integer | constraints.float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.49, 2.85))

	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.459, 2.85))
}
