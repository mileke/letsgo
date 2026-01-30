package main

import "fmt"

func main() {
	/*func (r receiver) identifier(p parameter) (return(s)) {
	code here
	}*/

	foo()
	bar("Mileke")
	fmt.Println(aloha("Big M"))
	fmt.Println(dogYears("Roni", 4))
}
func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years: "), age
}
