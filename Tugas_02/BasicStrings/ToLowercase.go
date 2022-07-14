package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.ToLower(os.Args[1]))
}

//https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/06-tolowercase/solution/main.go
