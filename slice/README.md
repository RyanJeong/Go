#   Slice<br>
##  배열은 크기가 자료형에 고정되어 있는 반면에 슬라이스는 길이와 용량을 갖고 있고 길이가 변할 수 있는 자료구조<br>

##  인덱스 세는 법<br>
### 시작 인덱스는 0부터 시작하고, 시작 인덱스에 있는 원소는 영역에 포함되나, 끝 인덱스에 있는 원소는 영역에 포함되지 않음<br>
```Go
for i := start; i < end; i++ {
	fmt.Println(A[i])
}
```
### 0부터 시작하고 끝 인덱스가 영역에 포함되지 않는 것은 여러 장점이 있음<br>
1. 길이를 구할 때 끝 인덱스에서 시작 인덱스를 빼면 구할 수 있으며, 불필요하게 +1 연산을 하지 않아도 된다.
```Go
length := end - start
```

2. 시작 인덱스와 길이를 알 경우, 끝 인덱스는 아래와 같이 단순하게 계산할 수 있다.
```Go
end := start + length
// 슬라이스의 구간: [start:start+length]
```

3. 영역 분할 시 편리함을 제공하는데, 예를 들어 0:n 구간을 m 등분 한다면 0:n/m, n/m:2*n/m, 2*n/m:3*n/m, ..., (m-1)*n/m:n과 같이 표현할 수 있다. 다시 설명하면 n개의 구간을 m등분하면 구간 하나 당 길이는 n/m이 되므로, 해당 구간 길이에 나누어진 구간의 인덱스를 곱해 구간을 맞추면 된다.<br>
```Go
package main

import "fmt"

func split(nums []int, m int) (result [][]int) {
	n := len(nums)
	for i := 0; i < m; i++ {
		result = append(result, nums[i*n/m:(i+1)*n/m])
	}
	return
}

func main() {
	fmt.Println(split([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
}
```

4. 두 구간이 서로 인접해 있는지 확인할 때 한 구간의 끝과 다른 구간의 시작 인덱스가 같은지 비교하면 된다. 

5. 연속된 영역에서 중간 부분을 제거하고자 할 때 편리함을 제공하는데, 예를 들어 15:25 구간에서 18:22 구간을 잘라내면 남는 부분은 15:18, 22:25가 된다. 이는 시작 인덱스에 있는 원소는 영역에 포함되고, 끝 인덱스에 있는 원소는 영역에 포함되지 않는다는 특징에 기인한다. 이 과정에서 불필요하게 +1 또는 -1 연산을 하지 않아 추가 연산이 필요 없고 이러한 연산 과정 중에 발생할 수 있는 오류도 예방할 수 있다.<br>

6. 직사각형 형태로 추상할 수 있는 2차원 배열을 1차원으로 표현할 때에도 편리함을 제공한다. 예를 들어, n X m 행렬이라고 할 때, n X M개의 1차원 배열이나 슬라이스를 잡은 뒤, i행 j열에 있는 원소는 A[i*m+j]로 찾으면 된다. 
```Go
package main

import "fmt"

func main() {
	n, m := 2, 3
	a := []int{
		1, 2, 3,
		4, 5, 6,
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%4d", a[i*m+j])
		}
		fmt.Println()
	}
}
```

```Go
// 빈 스트링을 n개 갖고 있는 슬라이스 만들기
slice := make([]string, n) // nil

// 값 0을 n개 갖고 있는 슬라이스 만들기
slice := make([]int, n) // 0
```
```Go
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
```

##  Go에서의 슬라이스는 음수를 사용할 수 없음
```Go
// 슬라이스 내에서 마지막 요소 값을 제외하고 싶을 때
// slice[:-1]
slice[:len(slice)-1]
```

##  슬라이스 덧붙이기
```Go
slice1 := append(slice, "Add")

// append 함수는 가변 인자를 수용함
slice2 := append(slice, "Add1", "Add2", "Add3")

// append 함수의 첫 파라미터는 슬라이스 객체이고, 두 번째는 추가할 요소의 값이기 때문에, 
// 아래와 같이 사용하면 두번째 파라미터가 값이 아닌 슬라이스이기 때문에 오류를 반환함
slice3 := append(s1, s2)

// 두번째 파라미터에 슬라이스의 요소들을 그대로 옮기고 싶을 경우, 아래와 같이 가변인자임을 표현할 것
slice4 := append(s1, s2...)
```
```Go
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
```

##  용량(Capacity)
### 뒤에 얼마나 덧붙일 공간이 있느냐를 기준으로 용량이 결정됨<br>
* 뒤에 요소 2개를 잘라낸 경우에는 길이는 2만큼 줄어들지만 여전히 2만큼 공간은 남아있는 상태이므로 용량은 5<br>
* 앞에 요소 2개를 잘라낸 경우에는 길이가 2만큼 줄어들고 뒤에 공간도 줄어든 상태이므로 용량도 3으로 줄어듦<br>
```Go
// 길이는 3, 용량은 5인 슬라이스를 만들고자 할 경우
nums := make([]int, 3, 5)

// 이는 아래 표현과 동일함
nums := make([]int, 5)
nums = nums[:3]

// 빈 슬라이스를 생성하되, 미리 공간을 예약해두고 싶은 경우
// 즉, n개까지 실이가 늘어나더라도 다른 곳으로의 메모리 복사를 방지하고자 할 경우
//   추가될 개수를 미리 예측할 수 있다면 이 방법은 매우 효율적인 방법임
nums := make([]int, 0, 15)
// nums = append(nums, x)를 수행해도 요소의 갯수가 15개를 넘지 않으면
// 다른 곳으로의 메모리 복사는 일어나지 않음
```
```Go
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
```

##  슬라이스의 내부 구현<br>
### 슬라이스는 마치 배열을 가리키고 있는 구조체와 같음<br>
*   배열의 시작 주소<br>
*   슬라이스의 길이<br>
*   슬라이스의 용량<br>
>   여러 슬라이스가 동일한 배열을 공유(참조)할 수 있는 이유<br>

```Go
nums = append(nums, 10)
// 1. len(nums) + 1 <= cap(nums): 용량을 초과하지 않을 경우
// 시작 위치로부터 길이만큼 오른쪽으로 이동한 위치에 새로운 값을 집어넣고,
// 길이가 증가한 슬라이스 반환(한 식에 슬라이스를 두 번 사용하는 이유)
// 2. len(nums + 1 > cal(nums): 용량을 초과하는 경우
// 더 큰 배열을 새로 하나 메모리 상에 할당하고, 해당 슬라이스도 바뀐 내용에 맞게
// 수정되어 반환됨
// 마찬가지로, 바뀐 슬라이스 내용을 받을 수 있도록 인자로 사용된 슬라이스에 다시 대치되어야 함
```
```Go
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
```
```Go
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
```
>   위의 두 예제를 보면, 슬라이스가 가리키는 배열의 크기보다 더 큰 값을 이어 붙여야 할 경우, 원본 배열의 크기의 정수배만큼 크기 할당을 하는 모습을 나타냄<br>
>   그럼, 아래 예제는 어떻게 동작할까?<br> 
```Go
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
```
>   처음 이어붙인 결과에서는, 원본 크기(2)의 정수배만큼 크키가 확장됐으며, 다음으로 이어붙인 결과에서는, 바뀐 원본 크기(10)의 정수배만큼 크기가 확장된 모습을 보였음<br>

##  슬라이스 복사<br>
```Go
func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
    // range 결과를 하나로 받으면 인덱스가 들어옴
    // range 결과를 두개로 받으면 첫 번째에는 인덱스, 두 번째에는 값이 들어옴
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)
	// Output:
	// [30 20 50 10 40]

	return
}
```
### 아래와 같이 간단히 사용할 수도 있음 <br>
```Go
copy(dest, src)

// 윗 표현을 보다 안정적으로 사용하고자 할 경우
// if문 내부에 ; 문자를 사용해 식을 구분할 수 있음
// 이처럼 표현하는 이유는! n은 if문 내에서 생성된 지역변수이므로,
// else-if문 밖으로 나갈 경우 해당 변수는 소멸됨
if n := copy(dest, src); n != len(src) {
    fmt.Println("복사가 덜 됐습니다.")
}

// dest의 크기를 애초에 src 크기와 동일하게 슬라이스를 할당하면 윗 문제를 해결할 수 있음
src := []int{10, 20, 30, 40, 50}
dest := make([]int, len(src))
copy(dest, src)

// append 함수를 사용해 간단히 복사된 슬라이스를 얻는 방법
src := []int{10, 20, 30, 40, 50}
dest := append([]int(nil), src...) // []int(nil): 빈 슬라이스

// 아래 식은?
dest := src // 슬라이스의 원소들이 복사되는 것이 아닌,
            // 슬라이스의 내용(배열 포인터, 길이, 용량)이 복사됨
            // 즉, src 원소 내용이 변경되면
            // dest 원소 내용도 같이 변경됨
```

##  슬라이스 삽입 및 삭제<br>
### 삽입<br>
```Go
// 첫 번째 방법:
// a 슬라이스의 i번째 원소로 x를 삽입하고자 할 경우:
a = append(a[:i+1], a[i:]...)
a[i] = x

// 설명: 
// [0] [1] [...] [i-1] [i]                      : a[:i+1]
//                         [i] [i+1] [...] [n-1]: a[i:]
// [0] [1] [...] [i-1] [i] [i] [i+1] [...] [n-1]: append(a[:i+1], a[i:]...)
// [0] [1] [...] [i-1]  x  [i] [i+1] [...] [n-1]: a[i] = x
// ---------------------------------------------------------------------------
// [i] 값을 두 번 복사한 후, i번째 위치한 곳에 x 값을 대입
// i가 0일 경우?
// [0] [0] ...
//  x  [0] ...
//
// i가 n-1일 경우?
// ... [n-1] [n-1]
// ... [n-1]   x
//

// i가 len(a)일 경우(즉, 총 길이가 n+1일 경우)
if i < len(a) {
    a = append(a[:i+1], a[i:]...)
    a[i] = x
} else {
    a = append(a, x)
}
```
```Go
// 두 번째 방법:
// a 슬라이스의 i번째 원소로 x를 삽입하고자 할 경우:

a = append(a, x)
copy(a[i+1:], a[i:])
a[i] = x

// 설명: 
// [0] [1] [...] [i-1] [i] [i+1] [...] [n-1]
// [0] [1] [...] [i-1] [i] [i+1] [...] [n-1] [x]   : a = append(a, x)
// [0] [1] [...] [i-1] [i] [i]   [i+1] [...] [n-1] : copy(a[i+1:], a[i:])
// [0] [1] [...] [i-1] [x] [i]   [i+1] [...] [n-1] : a[i] = x
//
// copy 연산을 할 때, 두 슬라이스 중에서 짧은 슬라이스 길이만큼 복사를 하므로,
// 맨 오른쪽 끝에 있는 x는 값 복사가 이루어지지 않음

// 붙이고자 하는 요소가 여러개일 때
x := []int{7, 8, 9}
a = append(a, x...)
copy(a[i+len(x):], a[i:])
copy(a[i:], x)
```

### 삭제<br>
```Go
// 요소 한 개 삭제
a = append(a[:i], a[i+1:]...)

// 일반적인 삭제 방법
a = append(a[:i], a[i+k:]...)

// 위의 방법들은 각 요소를 이동하면서 매번 복사 연산이 발생
// O(n)의 시간복잡도를 O(1)로 줄이는 방법은 아래와 같음
// (그러나, 순서는 바뀌게 된다!)
a[i] = a[len(a)-1]
a = a[:len(a)-1]

// 설명:
// [0] [1] [...] [i-1]  [x]  [i]   [i+1] [...] [n-1]
// [0] [1] [...] [i-1] [n-1] [i]   [i+1] [...] [n-1] : a[i] = a[len(a)-1]
// [0] [1] [...] [i-1] [n-1] [i]   [i+1] [...]       : a[i] = a[:len(a)-1]
```
```Go
// 여러 요소 삭제 시 예외상황
start := len(a)-k
if i+k > start {
    start = i+k
}
copy(a[i:i+k], a[start:])
a = a[:len(a)-k]

// Case #1. [0 1 2 3 4 5 6 7 8], i: 2, k: 3
//          [0 1 5 6 7 8]: OK
// Case #2. [0 1 2 3 4 5 6 7 8], i: 3, k: 4
//          [0 1 2 5 6 7 8]: Error!
// -> 삭제하고자 하는 영역의 바로 뒤부터 복사(O(n))
// -----------------------------------------------
//   [0 1 2 3 4 5 6 7 8]
//   [0 1 2 7 4 5 6 7 8] a[i] = a[i+k+1]
//   [0 1 2 7 8 5 6 7 8] a[i+1] = a[i+k+2]
//   [0 1 2 7 8]         a = a[:len(a)-k]
```
```Go
// 메모리 누수를 방지하는 방법
// 남아있는 길이는 nil로 일일히 지워주어야 함
// 1. 구조체 안에 있는 포인터들을 nil로 초기화
// 2. 해당 구조체를 빈 구조체(포인터가 모두 nil이고 다른 필드들도 기본값이 되는 구조체)
//    로 대치함

// 요소 하나를 삭제하는 경우
copy(a[:i], a[i+1:])
a[len(a)-1] = nil   // 생략 시 메모리 누수 위험
a = a[:len(a)-1]

// 요소 여러 개를 삭제하는 경우
copy(a[:i], a[i+1:])
for i := 0; i < k; i++ {
    a[len(a)-1-i] = nil // 생략 시 메모리 누수 위험
}
a = a[:len(a)-k]
```
### 포인터를 포함한 구조체의 경우에는, 위의 코드에서 nil 대신에 구조체 이름이 T일 경우에 T{}를 입력해 초기화하면 됨
