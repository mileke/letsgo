package main

import (
	"fmt"
	"io"
	"os"
)

//THis lecture walked through the importance of interfaces to type Writer. There's a lot to read from the standard library, but the code here just shows how these things can work interchangably cos of interfaces.

func main() {
	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground")
}
