package main

import "fmt"

func main() {
	x := 1
	if x == 1 {
		fmt.Println("vim-go")
	} else {
		fmt.Println("go")
	}

	i := 4
	if in := i * i; in == 16 {
		fmt.Println("go")
	}

	var msg string
	state := 2

	switch state {
	case 1:
		msg = "Good"
	case 2:
		msg = "So-so"
	case 3, 4:
		msg = "Bad!"
	default:
		msg = "Invalid"
	}
	fmt.Println(msg)

	switch new_state := state*2 - 5; new_state {
	case 1:
		msg = "Good"
	case 2:
		msg = "So-so"
	case 3, 4:
		msg = "Bad!"
	default:
		msg = "Invalid"
	}
	fmt.Println(msg)
}
