package main

import (
	"fmt"
)

func sliceaslice() {

	x := []int {4,5,7,8,42}
	fmt.Print(x[1])
	//colon allows you to slice the slice
	// so it will give you from position one to the end
	// so 5 7 8 42
	fmt.Print(x[1:])


	for i, v := range x {
		fmt.Println(i,v)
	}
	// 
	for i := 0, i <=4; i++ {
		fmt.Println(i, x[i])
	}
}