* 함수 리터럴(Function literal)은 함수 자체를 값으로 다룰 수 있는 상태를 나타냄 <br>
```Go
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
```
