package main

import (
	"fmt"
	"os"
)

func main() {
	// pass your name and lastname to the program

	fmt.Printf("Type of %.2f is %[1]T\n", 3.142)

	name, lastname := os.Args[1], os.Args[2]

	msg := "Your name is %s and your lastname is %s.\n"
	fmt.Printf(msg, name, lastname)
}
