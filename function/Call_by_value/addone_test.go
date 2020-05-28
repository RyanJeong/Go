package Add_One

import "fmt"

func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

func ExampleAddOne() {
	n := []int{2, 3, 4, 5, 6}
	AddOne(n)
	fmt.Println(n)
	// Output:
	// [3 4 5 6 7]

	return
}
