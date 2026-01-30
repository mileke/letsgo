package main

import "fmt"

func main() {
	defer foo()
	bar()
	rah()
}

func foo() {
	fmt.Println("foo")
}
func bar() {
	fmt.Println("bar")
}
func rah() {
	fmt.Println("rah")
}

// func (r receiver) identifier(p parameter) (returns) {code}
