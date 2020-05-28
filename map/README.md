#   Map<br>
*   Go에서 map은 해시테이블로 구현됨<br>
*   해시맵은 키와 값으로 구성되며, 키를 이용해 값을 상수 시간에 가져올 수 있음<br>
*   해시맵은 순서를 갖지 않음<br>

### Map 정의<br>
```Go
// map 정의, 그러나 아직 map 생성이 되지 않아 사용할 순 없는 상태
// nil 값을 가지고 있는 slice는 append로 덧붙일 수 있지만,
// map은 생성되어야 추가 가능
var m map[keyType]valueType

// map 생성
m := make(map[keyType]valueType)

// 빈 map으로 초기화
m := map[keyType]valueType{}
```

### Map 사용<br>
```Go
// map의 값 읽기
// 해당 키가 없으면, 자료형의 기본값 반환
// 정수형: 0, 문자형: 빈 문자열
m[key]  

// map의 값을 반환받을 때, 변수가 두 개일 경우:
// 첫 번째 변수: 값, 두 번째 변수: key 존재 여부
value, ok := m[key]

// map 값 쓰기
// key가 있을 경우 기존 값이 변경되고,
// key가 없을 경우 새로 생성됨
map[key] = value
```
### map을 사용하는 간단한 예:<br>
```Go
package count

func Count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}
```
```Go
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

    // map을 range로 받으면 키와 값이 나오며,
    // for _, v := range m { ... } 형태로 사용하면
    // key는 무시되고, 값만 받아올 수 있음
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
```

### Set<br>
* golang에서는 set을 정식으로 제공하지 않아 map을 응용하여 set과 비슷한 동작을 하도록 구현해서 사용할 수 있음<br>
```Go
// 1st version
// bool type이 메모리를 차지하는 단점 존재
package set

func hasDupeRune(s string) bool {
	runeSet := map[rune]bool{}
	for _, r := range s {
		if runeSet[r] {

			return true
		}
		runeSet[r] = true
	}

	return false
}
```
```Go
// 2st version
// 빈 구조체를 값으로 사용해 메모리 낭비를 방지
// struct{}처럼 아무 필드가 없는 구조체는 값이 아닌 자료형
// 해당 구조체에 아무 내용이 없는 상태의 값을 표현하려면?
// struct{}{}
package set

func hasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exists := runeSet[r]; exists {

			return true
		}
		runeSet[r] = struct{}{}
	}

	return false
}
```
```Go
package set

import "fmt"

func ExampleHasDupeRune() {
	fmt.Println(hasDupeRune("숨바꼭질"))
	fmt.Println(hasDupeRune("다시합시다"))
	// Output:
	// false
	// true

	return
}
```

* 삭제
```Go
// m[key]를 더 이상 사용할 수 없음
// len(m) 1 감소
delete(m, Key)
```

> The struct{} is a struct type with zero elements. It is often used when no information is to be stored. It has the benefit of being 0-sized, so usually no memory is required to store a value of type struct{}.<br>
> struct{}{} on the other hand is a composite literal, it constructs a value of type struct{}. A composite literal constructs values for types such as structs, arrays, maps and slices. Its syntax is the type followed by the elements in braces. Since the "empty" struct (struct{}) has no fields, the elements list is also empty:<br>
```Text
 struct{}  {}
|  ^     | ^
  type     empty element list
```

