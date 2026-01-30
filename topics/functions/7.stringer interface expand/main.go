package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("This is the book ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("The number is ", strconv.Itoa(int(c)))
}

func main() {
	b := book{
		title: "West With The Night",
	}

	var c count = 42

	log.Println(b)
	log.Println(c)
}
