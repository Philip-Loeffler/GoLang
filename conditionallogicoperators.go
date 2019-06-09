package main

import (
	"fmt"
)

func conditionaloperators() {
	// true and false, this will evaluate to true
	fmt.PrintLn(true && true)
	// this will be false, if it is a true and false it will be false
	fmt.PrintLn(true && false)
	// true or true? this will be true
	fmt.PrintLn(true || true)
	// true or false? this will be true
	fmt.PrintLn(true || false)
	// not true, this will be false
	fmt.Println(!true)
}
