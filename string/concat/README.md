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

### 문자열 이어 붙이기 벤치마크<br>

* N개의 문자열을 하나로 이어 붙이는 경우에는 일반적으로 bytes.Buffer를 이용하는 방법이 가장 빠른데, 이는 몇 개의 문자열을 이어붙이는가에 대해 컴파일 타임에 계산할 수 없거나, 만약 계산할 수 있다고 해도 여기까지 최적화 하지 않음<br>

* 그러나 단순히 크기가 정해진 숫자의 문자열들을 이어 붙일 경우에는 굳이 현란한 기술을 구사할 필요 없이 '+' 기호로 합치는 것이 가장 우수한 성능을 나타냄<br>

####  4개의 문자열을 다양한 방법을 통해 하나의 문자열로 합치는 예:<br>
##### 빠른 순으로 정렬되었음<br>

| Benchmark function | Repeat | Speed |
| --- | --- | --------- |
| BenchmarkPlus4-12    | 30000000 | 52.6 ns/op |
| BenchmarkJoin4-12    | 10000000 | 223 ns/op  |
| BenchmarkBytes-12    | 10000000 | 249 ns/op  |
| BenchmarkSprint4-12  |  3000000 | 561 ns/op  |
| BenchmarkSprintf4-12 |  2000000 | 630 ns/op  |

출처: https://github.com/jaeyeom/gogo

####  May 27, 2020, Results of the benchmark:<br>
##### 빠른 순으로 정렬되었음<br>
| Benchmark function | Repeat | Speed |
| --- | --- | --------- |
| BenchmarkPlus4-12    | 50000000 | 23.8 ns/op |
| BenchmarkJoin4-12    | 20000000 | 99.5 ns/op  |
| BenchmarkBytes-12    | 20000000 | 100 ns/op  |
| BenchmarkSprintf4-12 | 10000000 | 215 ns/op  |
| BenchmarkSprint4-12  | 10000000 | 222 ns/op  |

