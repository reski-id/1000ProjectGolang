package main

import (
	"fmt"
	"strings"
)

func main() {
	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))
}

//https://github.com/inancgumus/learngo/tree/master/08-numbers-and-strings/02-strings/exercises/07-trim-it
