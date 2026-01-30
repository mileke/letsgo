package main

import "fmt"

func main() {
	am := map[string]int{
		"Mileke": 23,
		"Steph":  43,
	}

	fmt.Println("The age of Mileke is:", am["Mileke"])

	an := make(map[string]int)

	an["Johnathan"] = 44
	an["Alexandra"] = 29

	fmt.Println("The age of Alexandra is:", an["Alexandra"])
	delete(am, "Mileke")
	fmt.Println(am)
	fmt.Println(len(an))

}
