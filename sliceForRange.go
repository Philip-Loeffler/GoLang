package main

import (
	"fmt"
)

func Range() {

	// if you're looping over an array, slice, string, or map
	// a rangle clause can manage the loop

	x := []int{4, 5, 7, 8, 42}
	// will tell you the length of the array
	fmt.Print(len(x))

	// will tell you the position.. so it would be 4
	fmt.Print(x[0])

	// this will loop over your array and tell you the index and the value
	for i, v := range x {
		fmtPrintln(i, v)
	}
}
