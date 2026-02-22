package main

import (
	"fmt"

	"github.com/mileke/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// fmt.Println(puppy.Bark())
	// fmt.Println(puppy.Barks())
}

/*func main() {

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}*/

/*func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! Welcome to the Go programming world.\n", name)

	//var for zero value
	var zero int
	fmt.Printf("The zero value of int is: %d\n", zero)

	//short delcaration operator
	short := 134
	fmt.Printf("The short value is: %d\n", short)

	//multiple variable initialisation
	x, y, z := 1, 2, 3
	fmt.Println("The values for multi initialisation are:", x, y, z)

	//specific type and value
	var w int = 42
	fmt.Printf("The specific value is : %d\nThe specific type is %T", w, w)

	//
	_ = name

}*/

/*func main() {

	name := "Mileke"
	age := 23
	pace := 5.5

	fmt.Printf("The customer's name is %v and of type %T", name, name)
	fmt.Printf("\nThe customer's age is %v and of type %T", age, age)
	fmt.Printf("\nThe customer's pace is %v and of type %T", pace, pace)

}
*/

/*func main() {
	var a int8 = 127
	var b uint8 = 255

	fmt.Printf("The value of a is %d and type is %T\n", a, a)
	fmt.Printf("The value of b is %d and type is %T\n", b, b)
}
*/
