package count

import (
	"fmt"
	"sort"
)

func ExampleCount() {
	codeCount := map[rune]int{}
	Count("가나다라마바사마다가", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
	// Output:
	// 가 2
	// 나 1
	// 다 2
	// 라 1
	// 마 2
	// 바 1
	// 사 1

	return
}
