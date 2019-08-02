package main

import (
	"fmt"
)

func functionSyntax() {
	foo()
	bar("james")
	phil("phil")
	// functions with return statements will have a variable associated 
	// with them
	s1 := woo("moneypenny")
	fmt.println(s1)
	x, y := mouse("ian", "flemming")
}

// func is a keyword
// when you define a function its called a parameter
// when you call a function and pass something in its called an argument
//func  (r receiver) identifiers(parameters)  (return(s)) {...}

//function that takes in no parameters
func foo() {
	fmt.Println("hello from foo")
}
//so s is the variable name and it is a string
// so when you call it it expects a string
// everything in go is passed in by VALUE
func bar(s string) {
	fmt.Println("hello" s)
}

func woo(st string) string {
	return fmt.Sprint("hello from woo,"st)
}
// will return a string and a bool
func(fn string, ln string) (string, bool)  {
a := fmt.Sprint(fn, ln, `, "hello`)
b := false
return a, b 
}

//ex

func phil(p string){
	fmt.Println("yo dawg dont trip", p)
}