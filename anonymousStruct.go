package main

import (
	"fmt"
)

// type person
// and first, last are values of type person
// explicit struct
type person struct {
	first string
	last  string
}

// this is is an anonomous because the struct doesnt have a name
func anonstruct() {
	//type
	p1 := struct {
		first string
		last  string
		age   int
	}{
		// values
		first: "james",
		last:  "bond",
		age:   32,
	}

	fmt.Println(p1)
}
