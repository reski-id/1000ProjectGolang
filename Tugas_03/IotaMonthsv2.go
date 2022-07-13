package main

import "fmt"

func main() {
	const (
		_   = iota // 0
		Jan        // 1
		Feb        // 2
		Mar        // 3
	)

	fmt.Println(Jan, Feb, Mar)
}

//https://github.com/inancgumus/learngo/tree/master/10-constants/exercises/08-iota-months-2
