// rememeber, control flow is sequencential, iterative, and imcremental
package main

import (
	"fmt"
)

func switchstatements() {

				switch {
				case false:
					fmt.Println("this should not print since its false")
				}
			case (2 == 4):
				fmt.Println("this should not print because 2 is not equal to 4")
			case (3 == 3 ):
				fmt.Println(" this will print since its true")
				fallthrough // if you do a fall through it will still dump
				// through the rest of the println
				// we generally do not even use this
				//because with case it will stop at the true statement
				// but fallthrough will just continue going 
				case (4 == 4)
				fmt.Println("this will print")
				default : // will defalt to this is nothing is true
				fmt.Println("this is default")

}