* Golang에서의 함수는 철저히 Call by Value를 따름
```Go
package Add_One

import "fmt"

func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

func ExampleAddOne() {
	n := []int{2, 3, 4, 5, 6}
    // 이때, AddOne의 nums []int 매개변수로 '값 복사' 발생
    // 결과값이 바뀐 이유는, 슬라이스 내 멤버변수(배열 포인터, 길이, 용량)
    // 복사가 일어난 상태에서 배열 포인터를 참조해 값을 수정했기 때문!
	AddOne(n)
	fmt.Println(n)
	// Output:
	// [3 4 5 6 7]

	return
}
```
