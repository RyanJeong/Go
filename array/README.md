#   Array<br>
##  배열은 연속된 메모리 공간을 순차적으로 이용하는 자료구조로써, 보통 배열을 단독으로 사용하기보다는 슬라이스를 같이 이용함<br>

```Go
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
```
