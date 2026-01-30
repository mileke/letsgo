package main

import (
	"fmt"
)

func main() {

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for i, v := range m {
		fmt.Printf("The key %v has the value %v\n", i, v)
	}
	fmt.Println(" ")
	age := m["Moneypenny"]
	fmt.Println("Moneypenny's age is", age)

	if v, ok := m["James"]; ok {
		fmt.Println("There is a lookup entry for James, and the age of Bond is:", v)
	}

	q := m["Q"]
	fmt.Println("The age of Q is", q)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int:", v)
	}
}
