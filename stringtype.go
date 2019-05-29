package main

import (
	"fmt"
)
// strings are immutable, so you cannot change the value stored, but you can assign a new value
// string type is a value that is stored in a variable of type string
// so you can assign a new value but you cannot change the type string 
func stringtype() {
	s := "string type"
	fmt.Println(s)
 // here we are assignning a new value 

	s = "new string type"
	fmt.Println("%T\n", s)

	//representing by numbers
	bs := []byte(s)
	fmt.Println(s)
	fmt.Println("%T\n", s)

	for i := 0, i <lens(s); i++ {
		fmt.Printf("%#U" , s[i])
	}
}
