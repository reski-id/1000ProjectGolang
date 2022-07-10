package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multi Assign #2
//
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//
//  2. Print the variables
//
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------

func main() {
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, temp, isTrue = "Mars", 19.5, true
	fmt.Println("Air is good on " + planet)
	fmt.Println("It's ", isTrue)
	fmt.Println(" It is", temp, "degrees")

	// ADD YOUR CODE BELOW
}

//https://github.com/inancgumus/learngo/tree/master/06-variables/04-assignment/exercises/06-multi-assign-2
