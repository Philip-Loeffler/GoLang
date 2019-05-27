package main

var a22 int

type hotdog int

var b22 hotdog

// type hotdog's underlying type is an intvar b hotdog

func main34() {
	a22 = 23
	b22 = 24
	fmt.println(a22)
	fmt.println(b22)
	// this expression int(b22) is converting variable b to a type
	// then assigning it to type a
	// this is called conversion in conversion, in other languages, it is called casting
	a22 = int(b22)
	// this will not work, because you cannot have something that
	// is type int be equal to something that is type "hotdog"
	a22 = b22
}
