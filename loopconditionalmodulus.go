package main

import (
	"fmt"
)


func loopconditionalmodulus() {
 // loop
	for i =: 0;  i < 100; i++ {
		// anytime any number is divided by three where the remainder is equal to zero
// print that number out

// condition with modulus 
				if i%3 == 0 {
					fmt.Println(i)
				}
	
	}