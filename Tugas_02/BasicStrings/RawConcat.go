package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Args[1]

	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)
}

//https://github.com/inancgumus/learngo/tree/master/08-numbers-and-strings/02-strings/exercises/03-raw-concat
