package main

import "fmt"

type person struct {
	first string
	age   int
}

func main() {
	p1 := person{
		first: "Mileke",
		age:   23,
	}
	p1.speak()
}

func (p person) speak() {
	fmt.Println("Your name is: ", p.first, " and your age is: ", p.age)
}
