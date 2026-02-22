package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucker gets filled."
	q := "Persistently, patiently, you are bound to suceed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("Value: %v \t Type: %T \t Data: %v\n", a, a, *a)
	fmt.Printf("Value: %v \t Type: %T \t Data: %v\n", b, b, *b)
	fmt.Printf("Value: %v \t Type: %T \t Data: %v\n", c, c, *c)
	fmt.Printf("Value: %v \t Type: %T \t Data: %v\n", d, d, *d)
}
