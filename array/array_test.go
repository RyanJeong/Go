package array

import (
	"fmt"

	"github.com/ryanjeong/go/string/hangul"
)

func Example_array() {
	//fruits := [3]string{"사과", "바나나", "토마토"}
	fruits := [...]string{"사과", "바나나", "토마토", "멜론"}

	for _, fruit := range fruits {
		//fmt.Printf("%t %s는 맛있다.\n", hangul.HasConsonantSuffix(fruit), fruit)
		fmt.Printf("%s", fruit)
		if hangul.HasConsonantSuffix(fruit) == true {
			fmt.Println("은 맛있다.")
		} else {
			fmt.Println("는 맛있다.")
		}
	}

	return
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 멜론은 맛있다.
}
