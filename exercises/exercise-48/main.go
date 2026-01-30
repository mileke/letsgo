package main

import (
	"fmt"
)

func main() {
	xxs := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "I'm 008."}}
	for _, v := range xxs {
		//fmt.Println(i, v)
		for y, x := range v {
			fmt.Println(y, x)
		}
	}
}
