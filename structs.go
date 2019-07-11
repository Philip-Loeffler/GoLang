// struct is a data structure
// allows us to compsoe together values of different types
// it aggregates together values of different types

package main

import (
	"fmt"
)
// the type is person
// the underlining type is struct 
type person struct {
	first string 
	last string 
}

func struct() {
// initalize each field on its on line
// we have created a value of type person
// create values of certain types.
// structs allow us to compose together the values of different types

	p1:= person{
		first: "james",
		last: "bond",
	}

	p2 := person{
		first: "miss",
		last: "moneypenny"
	}
	fmt.Println(p1)
	fmt.println(p2)

	fmt.println(p1.first)
	fmt.println(p2.last)
}


// structs is a sequence of named elements 

type person struct {
	// this is a named element in a field
	// each element had a name and a type 
	// first is name, second is type 
	// this is explicitily
	// and all of the names need to be unique
	first string 
	
}
type dude struct {
	// this is anonymous field, this is also an embedded field 
	 person
}