```Go
package array_vs_slice

import "fmt"

func Example_Comp_Arr_Slice() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [...]int{3, 4, 5, 6, 7}
	c := []int{5, 6, 7, 8, 9}

    fmt.Print("a ")
    printArr(a)
    //printSc(a)
    fmt.Print("b ")
    printArr(b)
    //printSc(b)
    //printArr(c)
    fmt.Print("c ")
    printSc(c)
    // Output:
    // a is an array
    // b is an array
    // c is a slice
}

func printArr(arr [5]int) {
    fmt.Println("is an array")

	return
}

func printSc(sc []int) {
    fmt.Println("is a slice")

	return
}
```
