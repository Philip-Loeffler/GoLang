package main

// a for statement specifies repated execution of a block

func forstatement() {
	//  repeats the execution of a block as long as a boolean condition evaluates to true

	// x < 10 is the condition
	// is 1 less than 10?
	// yes so then we increment it
	// then it comes back to the top as 2 and checks it again and continues the for loop
	x := 1
	for x < 10 {
		fmt.println(x)
		x++
	}

	fmt.println("done")

}
