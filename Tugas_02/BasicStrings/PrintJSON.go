package main

import "fmt"

func main() {
	// this one uses a raw string literal

	// can you see how readable it is?
	// compared to the previous one?

	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json)
}

//https://github.com/inancgumus/learngo/tree/master/08-numbers-and-strings/02-strings/exercises/02-print-json
