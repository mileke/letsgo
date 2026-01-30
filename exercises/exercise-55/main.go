package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors string
	color string
}

func main() {
	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Toyota",
		model: "Saloon",
		doors: "4-side open",
		color: "Silver",
	}

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Toyota",
		model: "Tacoma",
		doors: "2-side open",
		color: "Black",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println("A single field from v1 is: ", v1.make, v1.electric)
	fmt.Println("A single field from v2 is: ", v2.make, v2.electric)

}
