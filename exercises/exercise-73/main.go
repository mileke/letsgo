package main

import (
	"fmt"
	"time"
)

func main() {
	wrapper(wrapped)
}

func wrapper(wrapped func()) {
	startTime := time.Now
	wrapped()
	elapsedTime := time.Since(startTime())
	fmt.Println("The time that has elapsed is: ", elapsedTime)
}

func wrapped() {
	time.Sleep(2 * time.Second) // Simulate some work
	fmt.Println("MyFunction completed")

}
