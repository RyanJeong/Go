package array

import (
	"fmt"
)

func Example_array() {
	//fruits := [3]string{"사과", "바나나", "토마토"}
	fruits := [...]string{"사과", "바나나", "토마토"}

	for _, fruit := range fruits {
		fmt.Printf("%d %s는 맛있다.\n", hangul.HasConsonantSuffix(fruit), fruit)
	}

	return
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}
