// identifiers: an identifier is a sequece of one or more letters and digits
// the first character in an identer must be a letter.

// identifier = letter, letter | unicode_digit

// short declaration operator
// allows us to write code and declare variables.

package main

import (
	"fmt"
)

// := declares that what is before it is a variable and what is in front is going to be its value
// declare and assign
// if you declare a variable you have to use it

func fun() {
	x := 42
	fmt.Println(x)
	// assigning a new value
	x = 99
	fmt.Println(x)
	// this is a statement. statements are made up of expressions. 100 + 24 is the expression
	// statements usually occupy two lines
	y := 100 + 24
	fmt.Println(y)
}
