package main

import "fmt"

func main() {
	defer last()
	fmt.Println("This line should come first.")
	defer first()
	fmt.Println("This line should come third.")
}

func first() {
	fmt.Println("This is the first function that's deferred.")

}

func last() {
	fmt.Println("This is the last func to be deferred.")
}
