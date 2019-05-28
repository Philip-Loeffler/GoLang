package main

import (
	"fmt"
)

// vars in the package level scope
// declared variable x is of type int
var x int = 42
var y string = "james bond"
var z bool = "true"

// creating your own type
var a int

type burger int

var b burger

func main() {

	x := 42
	y := "james bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
