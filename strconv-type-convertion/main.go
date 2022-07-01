package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)

	fmt.Printf("%d, %s\n", i, s)
}
