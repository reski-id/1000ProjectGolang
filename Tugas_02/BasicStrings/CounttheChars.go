package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)
}

//https://github.com/inancgumus/learngo/tree/master/08-numbers-and-strings/02-strings/exercises/04-count-the-chars
