package array

import (
	"fmt"
)

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])
	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]

	return
}

func Example_append() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)     // f1 뒤에 f2 이어붙이기
	f4 := append(f1[:2], f2...) // "토마토" 제외하고 이어붙이기
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)
	// Output:
	// [사과 바나나 토마토]
	// [포도 딸기]
	// [사과 바나나 토마토 포도 딸기]
	// [사과 바나나 포도 딸기]

	return
}

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2))
	fmt.Println()

	sliced3 := sliced1[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	// 동일한 메모리를 참조하므로, 값이 모두 반영됨
	fmt.Println(nums, sliced1, sliced2, sliced3)

	// Output:
	// [1 2 3 4 5]
	// len: 5
	// cap: 5
	//
	// [1 2 3]
	// len: 3
	// cap: 5
	//
	// [3 4 5]
	// len: 3
	// cap: 3
	//
	// [1 2 3 4]
	// len: 4
	// cap: 5
	//
	// [1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]

	return
}

func Example_sliceCap2() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	nums = append(nums, 6)
	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	// 반환값을 받은 것에는 덧붙여진 슬라이스가,
	// 이전에 있던 슬라이스는 그대로(값은 바뀌었으나, 길이값은 변함 없으므로)
	slice := append(nums, 7)
	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	fmt.Println("slice:", slice)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice))
	fmt.Println()
	// Output:
	// nums: [1 2 3 4 5]
	// len: 5
	// cap: 5
	//
	// nums: [1 2 3 4 5 6]
	// len: 6
	// cap: 10
	//
	// nums: [1 2 3 4 5 6]
	// len: 6
	// cap: 10
	//
	// slice: [1 2 3 4 5 6 7]
	// len: 7
	// cap: 10

	return
}

func Example_sliceCap3() {
	nums := []int{1, 2, 3}

	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	nums = append(nums, 2)
	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	// 반환값을 받은 것에는 덧붙여진 슬라이스가,
	// 이전에 있던 슬라이스는 그대로(값은 바뀌었으나, 길이값은 변함 없으므로)
	slice := append(nums, 1)
	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	fmt.Println("slice:", slice)
	fmt.Println("len:", len(slice))
	fmt.Println("cap:", cap(slice))
	fmt.Println()
	// Output:
	// nums: [1 2 3]
	// len: 3
	// cap: 3
	//
	// nums: [1 2 3 2]
	// len: 4
	// cap: 6
	//
	// nums: [1 2 3 2]
	// len: 4
	// cap: 6
	//
	// slice: [1 2 3 2 1]
	// len: 5
	// cap: 6

	return
}

func Example_sliceCap4() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4, 5, 6, 7, 8, 9}
	nums3 := []int{0, 0, 0, 0}

	fmt.Println("nums1:", nums1)
	fmt.Println("len:", len(nums1))
	fmt.Println("cap:", cap(nums1))
	fmt.Println()

	nums1 = append(nums1, nums2...)
	fmt.Println("nums1:", nums1)
	fmt.Println("len:", len(nums1))
	fmt.Println("cap:", cap(nums1))
	fmt.Println()

	nums1 = append(nums1, nums3...)
	fmt.Println("nums1:", nums1)
	fmt.Println("len:", len(nums1))
	fmt.Println("cap:", cap(nums1))
	fmt.Println()
	// Output:
	// nums1: [1 2]
	// len: 2
	// cap: 2
	//
	// nums1: [1 2 3 4 5 6 7 8 9]
	// len: 9
	// cap: 10
	//
	// nums1: [1 2 3 4 5 6 7 8 9 0 0 0 0]
	// len: 13
	// cap: 20
	//

	return
}

func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)
	// Output:
	// [30 20 50 10 40]

	return
}
