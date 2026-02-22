package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)  // 42
	fmt.Println(&x) // address

	y := &x
	fmt.Printf("%v \t %T \n", y, y) //address and *int
	fmt.Println(*y)                 // 42
	fmt.Println(*&x)                // 42

	*y = 43
	fmt.Println(x)  // 43
	fmt.Println(*y) // 43
}
