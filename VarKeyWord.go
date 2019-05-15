package main

import (
	"fmt"
)

// if you need to declare variables outside of the function body
// also y is within the scope of the entire app
// where as x can only be used from where is its down.
var y = 43

// decalares there is a variable with the identify z, and the variable
// with the identify is type int. assigns the 0 value of type int to z
var z int

func main1() {
	// wont work because it is outside of the scope of x
	// fmt.Println(x)

	//declare a variable and assign a value (of a certain type)
	//short declaration works inside function body
	x := 42
	fmt.Println(x)

	fmt.Println(y)
	fool()
}
func fool() {
	fmt.Println("yo:", y)
}
