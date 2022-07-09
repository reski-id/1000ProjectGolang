package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//
//  Fix the code.
//
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be 32767
//
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

func main() {
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(int8(max) + min)
}

//https://github.com/inancgumus/learngo/tree/master/06-variables/05-type-conversion/exercises/05-convert-and-fix-5
