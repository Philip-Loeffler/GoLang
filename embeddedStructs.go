

package main

import (
	"fmt"
)

type person struct {
	first string 
	last string 
}
//taking type person, and embedding that into type secret agent
//because all secret agents are people
type secretAgent struct {
	person
	ltk bool 
}
func embeddedstruct() {

//outer type
	sa1:= secretAgent{
// person is the outer type
					person: person{
			first: "james",
			last: "bond",
			age: 32
		},
		ltk: true,
	}

	p2 := person{
		first: "miss",
		last: "moneypenny"
	}
	fmt.Println(p1)
	fmt.println(p2)

	fmt.Print(sa1.person.first)
	fmt.println(p1.first)
	fmt.println(p2.last)
}