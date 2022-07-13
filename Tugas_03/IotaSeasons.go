package main

import "fmt"

func main() {
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}

//https://github.com/inancgumus/learngo/tree/master/10-constants/exercises/09-iota-seasons/solution
