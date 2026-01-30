package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("My name is: ", p.first)
}

func main() {
	p1 := person{
		first: "Mileke",
	}

	p2 := person{
		first: "Toni",
	}

	p1.speak()
	p2.speak()
}
