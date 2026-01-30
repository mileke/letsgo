package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   24,
	}

	p2 := person{
		first: "Mileke",
		last:  "Kolawole",
		age:   23,
	}

	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "Chioma",
		last:  "Chizoba",
		age:   22,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("%#v\n", p3)

	fmt.Println("Mileke's age is:", p2.age)
}
