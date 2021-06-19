package main

import "fmt"

func main() {
	a := 0x0A // 1010
	b := 0x07 // 0111
	fmt.Println(a &^ b)
}
