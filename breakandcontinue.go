package main

import (
	"fmt"
)

//break will "break" you out of the loop
func nestingloop() {
	//code runs and checks if 1 is greater than 9. see's that its not so it loops till 10.
	// once it gets to ten it "breaks" off and stops
	x := 0
	for {
		x++
		if x > 100 {
			// if we hit something greater than 100 we are done
			break
		}
		// if x modules 2 is not equal to zero
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}

}
