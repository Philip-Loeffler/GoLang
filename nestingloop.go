package main

import (
	"fmt"
)

// outer loop is going to run as many times as we ask it to run
// so if the outer loop runs 10 times, and the inner loop runs five times
// the outer loop will run once, then the inner loop five times
// then outer loop runs once again, once the five times is over
func nestingloop() {

for i := 0 i <=10; i++ {
	for j := 0 j <3; j++ {

	fmt.Println(i, j)
	}

}

}
