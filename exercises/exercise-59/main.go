package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers)
	fmt.Println(foo(numbers...))

	fmt.Println(numbers)
	fmt.Println(bar(numbers))
}

func foo(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func bar(b []int) int {
	total := 0
	for _, v := range b {
		total += v
	}
	return total
}
