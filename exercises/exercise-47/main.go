package main

import (
	"fmt"
)

func main() {
	//create an array which holds 5 values of type int
	xs := make([]string, 0, 50)
	xs = append(xs, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`)
	fmt.Println(xs)
	fmt.Println(len(xs))
	fmt.Println(cap(xs))
	for i := 0; i < 50; i++ {
		fmt.Printf("The value of the slice at position %v is: %v and the type is %T\n", i, xs[i], xs[i])
	}
}
