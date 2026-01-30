package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("The anonymous func ran!")
	}()

	func(s string) {
		fmt.Println("This is an anon func showing my name", s)
	}("Mileke")
}

func foo() {
	fmt.Println("This is foo!")
}
