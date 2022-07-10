package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

func main() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}

//https://github.com/inancgumus/learngo/tree/master/06-variables/05-type-conversion/exercises/02-convert-and-fix-2
