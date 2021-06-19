package main

import "fmt"

func main() {
	a := 'a'
	str := "Hello\n"
	str_raw := `Hello\nWorld`
	fmt.Printf("%c", a)
	fmt.Println(str)
	fmt.Println(str_raw)
	fmt.Println(str_raw)
	fmt.Println(str_raw)

	line1 := `Hello
World`
	line2 := "Hello\n" + "World"
	fmt.Println(line1)
	fmt.Println(line2)
}
