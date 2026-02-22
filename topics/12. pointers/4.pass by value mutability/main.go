package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func mapDelta(m map[string]int, s string) {
	m[s] = 77
}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	m := make(map[string]int)
	m["James"] = 42
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])
}
