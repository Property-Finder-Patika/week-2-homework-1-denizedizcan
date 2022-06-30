package main

import (
	"fmt"
)

func main() {
	i := 314
	f := 3.14
	s := "Hello"
	b := true

	fmt.Println(
		"i:", i,
		"f:", f,
		"s:", s,
		"b:", b,
	)
	//multiple declaration
	d, e := 14, true

	fmt.Println(d, e)

	//short with expression
	sum := 27 + 3.5

	fmt.Println(sum)

	//discard values in a short declaration

	on, _ := true, true

	fmt.Println(on)

	// redeclare

	age, yourAge := 10, 20
	age, ratio := 42, 3.14

	fmt.Println(age, yourAge, ratio)
}
