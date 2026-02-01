package main

import "fmt"

func main() {
	x := 8
	y := &x
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*y)
}
