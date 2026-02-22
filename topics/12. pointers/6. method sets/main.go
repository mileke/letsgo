package main

import "fmt"

type dog struct {
	first string
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running")
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	d1.run()
	youngRun(d1)
}
