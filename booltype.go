package main

import (
	"fmt"
)

// boolean is nothing more than true or false.
// x is of type bool. x holds values of type bool.
var x bool

func main() {

	a := 7
	b := 42
	// does a equal b? this will evaluate to false
	fmt.Print(a == b)
	// not equal to
	fmt.print(a != b)
	// there is no triple equal sign in go
	// this one will be false
	fmt.Printf(x)
	x = true
	// this one will be true
	fmt.Print(y)
}
