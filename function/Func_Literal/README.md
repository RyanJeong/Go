package Func_Literal

import "fmt"

func Example_funcLiteral() {
	func() {
		fmt.Println("Hello!")
	}()
	// Output:
	// Hello!

	return
}

func Example_funcLiteralVar() {
	printHello := func() {
		fmt.Println("Hello!!")
	}
	printHello()
	// Output:
	// Hello!!

	return
}
