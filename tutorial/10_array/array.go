package main

import "fmt"

func main() {
	/* -------- array -------- */
	// array: 'specifies a length' and an element type
	// The type [n]T is an array of n values of type T
	// ex. type [2]int is defferent than type [4]int

	var a1 [3]int // declaration
	// By default an array is zero-valued
	fmt.Println(a1)
	a1[0] = 1
	a1[1] = 2
	a1[2] = 4
	fmt.Println(a1)
	var a2 = [3]int{1, 2, 4} // declare and initialize
	fmt.Println(a2)
	var a3 = [...]int{1, 2, 4}
	fmt.Println(a3)
	var a4 [2][3]int
	a4[0][0] = 5
	a4[0][1] = 4
	a4[0][2] = 3
	a4[1][0] = 2
	a4[1][1] = 1
	a4[1][2] = 0
	fmt.Println(a4)
	var a5 = [2][3]int{
		{5, 4, 3},
		{2, 1, 0}, // MUST have comma
		// if it wouldn’t be there, Go would have added a semicolon(;)
		// by the rule which would have CRASHED the program
	}
	fmt.Println(a5)

}
