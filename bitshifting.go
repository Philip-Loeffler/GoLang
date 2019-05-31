package main

import (
	"fmt"
)

func bitshiftingc() {
	x := 4
	fmt.Printf("%d\t\t%b", x, x)

	// y is eqaul to x but shift that over one
	y := x << 1
	fmt.Printf("%d\t\t%b", y, x)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

}
