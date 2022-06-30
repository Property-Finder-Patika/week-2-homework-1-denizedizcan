package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))

	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++

	factor--
	factor--

	fmt.Println(float64(counter) * factor)

	//circle area
	var (
		radius = 10.
		area   float64
	)

	area = math.Pi * radius * radius

	//sphere volume
	fmt.Printf("radius: %g -> area: %.2f\n",
		radius, area)

	vol := (4 * math.Pi * math.Pow(radius, 3)) / 3

	fmt.Printf("radius: %g -> volume: %.2f\n", radius, vol)

}
