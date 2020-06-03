package main

import "fmt"

func ExampleFunctions_print() {
	list1 := [...]int{1, 2, 3}
	list2 := [...]int{4, 5, 6}
	str1 := "STR1"
	str2 := "STR2"

	fmt.Println(list1, list2)
	fmt.Println(str1, str2)
	fmt.Print(list1, list2)
	fmt.Print(str1, str2)
	fmt.Printf("%d%d", list1, list2)
	fmt.Printf("%s%s", str1, str2)
	fmt.Printf("%d %d", list1, list2)
	fmt.Printf("%s %s", str1, str2)

	// Output:
	// [1 2 3] [4 5 6]
	// [1 2 3] [4 5 6][1 2 3][4 5 6][1 2 3] [4 5 6]

	return
}
