package main

import "fmt"

func main() {
	/*func (r receiver) identifier(p parameter) (return(s)) {
	code here
	}*/

	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum of the numbers is:", sum)
}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	sum := 0
	for _, v := range ii {
		sum += v
	}
	return sum

}
