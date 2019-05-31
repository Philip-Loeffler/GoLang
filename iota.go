package main

import (
	"fmt"
)

// iota is a pre declared identify

const (
	a = iota
	b = iota
	c = iota
)
const (
	d = iota
	e = iota
	f = iota
)

func iota() {
	// if you run this it will increment by one. so a = 0 b = 1 c = 2.
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
	// the second const will restart the increment so you will see 0 1 2
	fmt.Print(d)
	fmt.Print(e)
	fmt.Print(f)

}
