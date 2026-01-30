package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return rand.Intn(10)
}
func bar() (int, string) {
	wineAge := rand.Intn(23)
	wine := "Casamigos"
	return wineAge, fmt.Sprint(" is the possible age of wine ", wine)

}
