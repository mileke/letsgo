package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a()
}

var a = func() {
	fmt.Println(rand.Intn(7))
}
