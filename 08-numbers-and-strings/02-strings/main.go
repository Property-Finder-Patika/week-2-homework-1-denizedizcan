package main

import (
	"fmt"
	"strings"
)

func main() {
	// this one uses a raw string literal

	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json)

	// tolower

	fmt.Println(strings.ToLower("DeNiZ"))

	// trim string

	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))
}
