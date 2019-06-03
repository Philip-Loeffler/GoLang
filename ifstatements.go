package main

import (
	"fmt"
)

func ifstatement() {

// this will print everything because of the scope of x. it is declared 
// about the brackets so it follows all the way down
 x:= 42 
 if x ==42 {
	 fmt.Println("nice")
 }
 fmt.Println(x)
}

func scope() {
	 // this will not print out the x because it is out of the scope and 
	 // cannot see the value of x
	if x:= 42; x ==42 {
		fmt.Println("nice")
	}
	fmt.Println(x)
   }







	if x := 42; x == 42 {
		fmt.println("yo")
}
// true and false are pre declared constants
	if true {
			fmt.Println("001")
	}

	if false {
		fmt.Println("002")
}
// if its NOT true meaning its false it wil run...so this wont run 
if  !true {
	fmt.Println("003")
}
// if its NOT false meaning its true it will run so this will run
if !false {
fmt.Println("004")
}
 
if 2! = 2 {
	fmt.println("005")
}
if 2! = 2 {
	fmt.println("006")
}
// here is an itialize statement
// will run because it is true, x is equal to 42
if x := 42; x == 42 {
		fmt.println("yo")
}
	// wil not run because x is not equal to 43
if x := 42; x == 43 {
	fmt.println("yo")
}
