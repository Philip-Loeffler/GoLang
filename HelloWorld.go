package main

import "fmt"

// control flow, the order in which individual statements are executed or called
// ex: 1.sequence 2. loop:iterate 3.conditional

func main() {
	fmt.Println("hello world")

	foo()

	fmt.Println("here is some more stuff")

	// and here is a for loop
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)

		}

	}
}

// creating a reusable function
func foo() {
	fmt.Println("im in a foo man im in a foo")
}
