package main

import "fmt"

func addOne(v int) int {
	return v + 1
}

func pointerAddOne(v *int) {
	*v += 1
}

func main() {
	a := 1

	fmt.Println(a)         //1
	fmt.Println(addOne(a)) //2
	fmt.Println(a)         //1
	pointerAddOne(&a)
	fmt.Println(a) //2

}
