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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.fav {
			fmt.Println(v.first, v.last, v2)
		}
	}
}
