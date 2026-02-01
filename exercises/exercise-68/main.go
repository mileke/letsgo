package main

import "fmt"

func main() {
	fmt.Println(func(a int) int { return a }(12))
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println("The second func is:", i)
		}
	}()
}
