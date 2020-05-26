package concat

import (
	"fmt"
	"strings"
)

func Example_strCat() {
	// 문자열은 읽기 전용
	// s  := "abc"
	// ps := &s
	// s  += "def"
	// fmt.Println(s)
	// fmt/Println(*ps)
	// ------
	// ps는 s의 주소를 가리키는 포인터이며,
	// s 문자열에 "def"를 이어붙이게 되면
	// s 문자열에 "def"가 붙는 것이 아닌 새로운 문자열이 생성된다.
	// 또한, s의 주소도 변경되는데, "abc"를 가리키는 기존 주소에서
	// "abcdef" 문자열이 새로 생성된 주소를 가리킨다.
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)

	s1 := "ABC"
	s1 = fmt.Sprint(s1, "DE")
	s2 := "FGH"
	s2 = fmt.Sprintf("%sIJ", s2)
	s3 := "KL"
	s3 = strings.Join([]string{s3, "MN"}, "")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	// Output:
	// abcdef
	// abcdef
	// ABCDE
	// FGHIJ
	// KLMN
}
