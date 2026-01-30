package main

import "fmt"

func main() {
	random := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Mileke",
		friends: map[string]int{
			"Cole":  4,
			"Jenny": 7,
		},
		favDrinks: []string{"Sprite", "Chivita"},
	}

	fmt.Println("The full detail is:", random)
	fmt.Println("The person's name is:", random.first)
	fmt.Println("The person's name is:", random.friends["Cole"])
	fmt.Println("The person's name is:", random.friends["Cole"])

}
