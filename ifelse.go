package main

import (
	"fmt"
)


func ifelse() {
		x:=42
	if  x == 42 {
		fmt.println("our value was 40")
	} else {
		fmt.Println("our value is not 40")
}

// checks the first if, no dice, then checks the second one, still no dice, 
// then goes to the third one and bam snacks on snacks 
if  x == 42 {
	fmt.println("our value was 40")
} else if  x == 41{
	fmt.Println("our value is not 41")
} else  x ==42 {
	fmt.Println("our value is not 42")
}