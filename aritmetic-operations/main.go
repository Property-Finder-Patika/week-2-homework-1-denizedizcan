package main

import (
	"fmt"
	"math"
)

func main() {

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
