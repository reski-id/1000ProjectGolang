package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multi Short Func
//
// 	1. Declare two variables using multiple short declaration syntax
//
//  2. Initialize the variables using `multi` function below
//
//  3. Discard the 1st variable's value in the declaration
//
//  4. Print only the 2nd variable
//
// NOTE
//  You should use `multi` function
//  while initializing the variables
//
// EXPECTED OUTPUT
//  4
// ---------------------------------------------------------

func main() {
	// ADD YOUR DECLARATIONS HERE
	a, b := 4, 4
	b = a

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(b)
}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}

//https://github.com/inancgumus/learngo/tree/master/06-variables/04-assignment/exercises/07-multi-short-func
