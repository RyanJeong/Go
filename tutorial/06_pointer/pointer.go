package main

import "fmt"

func main() {
	var p *int
	fmt.Println(p) // <nil>

	i := 42
	p = &i // p points to i

	fmt.Println(*p) // 42, read i through the pointer
	fmt.Println(p)  // 0xc0000a4018, print i's address
	*p = 21         // dereferencing or indirecting
	// set i through the pointer
	fmt.Println(*p) // 21, read the new value of i
	fmt.Println(p)  // 0xc0000a4018
}
