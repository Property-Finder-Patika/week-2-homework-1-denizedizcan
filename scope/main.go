package main

import (
	"fmt"
)

func main() {

	x := "deniz"

	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}
