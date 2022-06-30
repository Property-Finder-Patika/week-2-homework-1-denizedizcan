package main

import "fmt"

func main() {
	color := "green"

	color = "blue"

	fmt.Println(color)

	color = "dark " + color

	fmt.Println(color)

	n := 0.

	n = 3.14 * 2

	fmt.Println(n)

	var (
		perimeter     int
		width, height = 5, 6
	)

	// first calculates: (width + height)
	// then            :  multiplies it with 2

	perimeter = 2 * (width + height)

	fmt.Println(perimeter)

	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	fmt.Println(lang, "version", version)

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5

	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp, "degrees")

	_, b := multi()

	fmt.Println(b)

	//swapping

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)
}

func multi() (int, int) {
	return 5, 4
}
