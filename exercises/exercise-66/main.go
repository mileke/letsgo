package main

import "math/rand"

func main() {
	Rando(40)
}

// Rando is a function that generates random numbers
func Rando(a int) int {
	x := rand.Intn(a)
	return x
}
