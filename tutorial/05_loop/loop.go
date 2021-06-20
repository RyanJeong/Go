package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	colors := []string{"Red", "Green", "Blue"}

	for index, color := range colors {
		fmt.Println(index, color) // 0 Red
		// 1 Green
		// 2 Blue
	}

	// You can skip the index or value by assigning to _.
	for _, color := range colors {
		fmt.Println(color) // Red
		// Green
		// Blue
	}

	a := 0
	for a < 10 {
		if a == 1 {
			a += 4
			continue
		}
		a++
		if a > 5 {
			break
		}
	}
	if a == 6 {
		goto JUMP
	}
	fmt.Println(a) // Can't access this line
JUMP:
	fmt.Println("JUMP") // JUMP

IGNORE_LOOP1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			break IGNORE_LOOP1
		}
	}
	// i=0 j=0

IGNORE_LOOP2:
	switch 1 {
	case 1:
		fmt.Println(1)
		for i := 0; i < 10; i++ {
			break IGNORE_LOOP2
		}
		fmt.Println(2)
	}
	fmt.Println(3)
	// 1
	// 3
	test()
}

func test() {
IGNORE_LOOP:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
		}
	}

	switch 1 {
	case 1:
		fmt.Println(1)
		for i := 0; i < 10; i++ {
			break IGNORE_LOOP
		}
		fmt.Println(2)
	}
	fmt.Println(3)

}
