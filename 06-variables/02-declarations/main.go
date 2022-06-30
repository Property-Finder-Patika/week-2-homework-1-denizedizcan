package main

import (
	"fmt"
)

var isLiquid bool

func main() {
	//int
	var height int
	fmt.Println(height)
	//bool
	var isOn bool
	fmt.Println(isOn)
	//float64
	var brightness float64
	fmt.Println(brightness)
	//strings
	var s string
	var r rune  // also a numeric type
	var by byte // also a numeric type
	fmt.Printf("s (%T): %q\n", s, s)
	// complex types
	var c64 complex64
	var c128 complex128

	// string types

	fmt.Println(

		c64, c128,
		r, by,
	)

	//multiple
	var (
		active bool
		delta  int
	)
	fmt.Println(active, delta)

	var firstName, lastName string = "", ""
	fmt.Printf("%q %q\n", firstName, lastName)

	//unused
	var isLiquid bool
	_ = isLiquid

}
