package main

import (
	"fmt"
)

var a = 42

//declare the variable with the identifier z
// it is of type string
// and i assign the value "y dawg how you do"
// z will only hold values of strings
// this is a static programing language
// a variable is declared to hold a value of a certain type
var x string = "yo dawg how you do"

func main4() {
	fmt.Println(a)
	// when you run this the complier will tell you what type the variable is
	fmt.Printf("%T\n", a)
}

// primitive data types
// bool, string

// composite/ agregate
// slice, struct
