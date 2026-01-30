package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini"}
	jmp := []string{"Jenny", "Moneypenny", "Fanta"}
	fmt.Println(jb, jmp)

	xxp := [][]string{jb, jmp}
	fmt.Println(xxp)
}
