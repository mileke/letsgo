package main

import (
	"fmt"
)

func main() {
	mp := map[string][]string{
		"Bond_James":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	mp["fleming_ian"] = []string{"steaks", "cigars", "espionage", "trials"}

	delete(mp, "no_dr")

	// fmt.Println(mp["Bond_James"])
	for k, v := range mp {
		fmt.Println(k)
		for i, v := range v {
			fmt.Printf("Index: %v \t slice value: %v\n", i, v)
		}
	}
}
