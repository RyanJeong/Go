package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {
	// New list.
	values := list.New()
	// Add 3 elements to the list.
	values.PushBack("bird")
	values.PushBack("cat")
	values.PushFront("snake")
	// Add 100 elements at the front.
	for i := 0; i < 20; i++ {
		// Convert ints to strings.
		values.PushFront(strconv.Itoa(i))
	}

	// Loop over container list.
	for temp := values.Front(); temp != nil; temp = temp.Next() {
		fmt.Println(temp.Value)
	}
}
