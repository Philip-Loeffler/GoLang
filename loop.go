package main

import (
	"fmt"
)

func loop() {

	// for init; condition; post. this is the structure of a loop
	// while i is less than or equal to a 100 it is going to run through this loop
	// i want you to add one to i.
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
