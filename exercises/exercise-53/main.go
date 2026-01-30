package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	p1 := person{
		first: "Mileke",
		last:  "Kolawole",
		fav:   []string{"Vanilla", "Strawberry"},
	}

	p2 := person{
		first: "Chioma",
		last:  "Chizoba",
		fav:   []string{"Chocolate", "Banana"},
	}

	for _, v := range p1.fav {
		fmt.Println(p1.first, "favorite Ice cream is: ", v)
	}
	for _, v := range p2.fav {
		fmt.Println(p2.first, "favorite Ice cream is: ", v)
	}
}
