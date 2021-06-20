package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Before", i) // 1
	val(i)
	fmt.Println("After ", i) // 1
	fmt.Println("Before", i) // 1
	ref(&i)
	fmt.Println("After ", i) // 2

	say("Hello", "World", "Golang") // Hello
	// World
	// Golang
	say("Variadic Function") // Variadic Function

	fmt.Println(calc("Rectangle", 10, 20)) // 200
	fmt.Println(calc("Square", 5))         // 25

	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	result := sum(1, 2, 4, 8, 16, 32) // 63
	fmt.Println(result)
}

func val(i int) {
	i++
	return
}

func ref(i *int) {
	*i++
	return
}

func say(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func calc(str string, y ...int) int {
	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}

	return area
}
