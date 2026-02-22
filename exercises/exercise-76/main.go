package main

import "fmt"

type dog struct {
	first string
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm walking")
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1)
	d1 = d1.changeName("Rover")
	youngRun(d1)
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}
