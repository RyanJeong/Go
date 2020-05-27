```Go
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
```

```Text
$ go test
PASS
ok
```

** 문자열 이어 붙이기 벤치마크

N개의 문자열을 하나로 이어 붙이는 경우에는 일반적으로 bytes.Buffer를 이용하는
방법이 가장 빠릅니다. 이것은 몇 개의 문자열을 이어붙이는지 컴파일 타임에 알 수
없거나 알 수 있다고 해도 거기까지 최적화를 하지는 않기 때문입니다.

그러나 단순히 정해진 숫자의 문자열들을 이어 붙일 때에는, 굳이 현란한 테크닉을
구사할 필요가 없습니다. 그냥 + 기호로 합치는 것이 가장 빠릅니다.

다음은 4개의 문자열을 다양한 방법을 이용하여 하나로 합친
경우입니다. 빠른 것이 위에 오도록 정렬하였습니다.

| BenchmarkPlus4-12    | 30000000 | 52.6 ns/op |
| BenchmarkJoin4-12    | 10000000 | 223 ns/op  |
| BenchmarkBytes-12    | 10000000 | 249 ns/op  |
| BenchmarkSprint4-12  |  3000000 | 561 ns/op  |
| BenchmarkSprintf4-12 |  2000000 | 630 ns/op  |

출처: https://github.com/jaeyeom/gogo

** May 27, 2020, Results of the benchmark:
| BenchmarkPlus4-12    | 50000000 | 23.8 ns/op |
| BenchmarkJoin4-12    | 20000000 | 99.5 ns/op  |
| BenchmarkBytes-12    | 20000000 | 100 ns/op  |
| BenchmarkSprintf4-12 | 10000000 | 215 ns/op  |
| BenchmarkSprint4-12  | 10000000 | 222 ns/op  |

