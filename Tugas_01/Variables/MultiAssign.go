package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multi Assign
//
//  1. Assign "go" to `lang` variable
//     and assign 2 to `version` variable
//     using a multiple assignment statement
//
//  2. Print the variables
//
// EXPECTED OUTPUT
//  go version 2
// ---------------------------------------------------------

func main() {
	// DO NOT TOUCH THIS
	var (
		lang    string
		version int
	)

	lang, version = "go", 2
	// ADD YOUR CODE BELOW

	// DO NOT TOUCH THIS
	fmt.Println(lang, "version", version)
}

//https://github.com/inancgumus/learngo/tree/master/06-variables/04-assignment/exercises/05-multi-assign
