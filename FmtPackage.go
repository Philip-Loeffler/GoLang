package main

import (
	"fmt"
)

var x1 = 42

func main111() {

	fmt.Printf("%T\n", x1)
	// will print binary
	fmt.Printf("%b\n", x1)
	// will run hexidecimal
	fmt.Printf("%x\n", x1)

}
