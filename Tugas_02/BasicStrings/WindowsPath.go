package main

import "fmt"

func main() {
	// this one uses a raw string literal

	// can you see how readable it is?
	// compared to the previous one?

	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`

	fmt.Println(path)
}

//https://github.com/inancgumus/learngo/tree/master/08-numbers-and-strings/02-strings/exercises/01-windows-path
