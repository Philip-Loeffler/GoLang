//grouping data together.

package main

import (
	"fmt"
)

func array() {
	// an array of type int, five is the size
	var x [5]int
	fmt.Print(x)
	x[3] = 42
	// i want to put 42 in position three
	// will be at position four
	fmt.Println(x)
	// this is just going to telll the length of the array
	fmt.Println(len(x))
	//
}
