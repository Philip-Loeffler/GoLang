package main

func deletingfromaslice() {

	x := []int{4, 5, 7, 8, 42}

	// append to slice position four
	// will pull out 7 and 8
	x = append(x[:2], x[4:]...)

}
