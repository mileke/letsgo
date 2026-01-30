package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := range 100 {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("The value of x is 3: %v and this has run %v times.\n", x, i)
		}

	}
}
