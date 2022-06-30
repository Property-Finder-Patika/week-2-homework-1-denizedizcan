package main

import "fmt"

func main() {
	const (
		home   = "Cemre Ozeren Deniz Edizcan"
		length = len(home)
	)

	fmt.Printf("There are %d characters inside %q\n",
		length, home)

	// Make them typeless
	const (
		width  = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
}
