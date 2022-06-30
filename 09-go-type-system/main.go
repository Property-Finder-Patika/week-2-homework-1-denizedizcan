package main

import "fmt"

func main() {
	// an english letter (search web for: ascii control code)
	var letter byte = 'A'
	fmt.Println("an english letter:", letter)

	// a non-english letter (search web for: unicode codepoint)
	var unicode rune = 'Ç'
	fmt.Println("a non-english letter:", unicode)

	// a year in 4 digits like 2040
	var year uint16 = 2040
	fmt.Println("a year in 4 digits like 2040:", year)

	// a month in 2 digits: 1 to 12
	var month uint8 = 6
	fmt.Println("a month in 2 digits: 1 to 12:", month)

	// the speed of the light
	var lightSpeed uint32 = 670616629 // miles
	fmt.Println("the speed of the light:", lightSpeed)

	// angle of a circle
	var angle float32 = 35.8
	fmt.Println("angle of a circle:", angle)

	var pi float64 = 3.141592653589793
	fmt.Println("PI number:", pi)

	var (
		width  uint16
		height uint16
	)

	width, height = 255, 265
	width += 10

	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("are they equal?", width == height)

	//convert the types

	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	fmt.Println(celsius, fahr)
}
