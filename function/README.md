#   Function<br>
##  Go 언어에서의 함수는 철저한 Call by Value를 따름<br>
```Go
func AddOne(nums []int) {
    for i := range nums {
        nums[i]++
    }
}

func ExampleAddOne() {
    n := []int{1, 2, 3}
    AddOne(n) // slice 정보가 복사되고,
              // AddOne 함수에서 값을 수정하면
              // ExampleAddOne 함수 내부에서도 값이 수정되는데,
              // 이는 슬라이스 내부 변수(배열 포인터, 길이, 용량)들이 복사되어 
              // AddOne 함수 내부에서 포인터에 의한 참조 연산을 수행하기 때문
    fmt.Println(n)
    // Output:
    // [2 3 4]
}
```

##  반환값이 여러 개일 수 있음<br>
*   예를 들어, C 언어에서는 함수의 반환값이 한 개이므로, 에러 값을 전달하기 위해 여러 어색한 방법을 사용해야 했음<br>
##  반환값들 중에서 에러를 반환할 경우, 보통 마지막 반환값에 에러 정보를 담는 것이 관레<br>
### 보통 일반적인 흐름에서 처리할 수 있는 오류 같은 경우는 에러 상황으로 간주하고, 심각한 오류를 야기하는 경우는 패닉(panic) 상황으로 간주해 처리함<br>
```Go
func WriteTo(w io.Writer, lines []string) (int64, error)

// 함수 호출 후 반환되는 예외값을 바로 처리할 수 있는 패턴
if err := MyFunc(); err != nil {    // if문 벗어나면 err 할당 해제

}

if _, err := MyFunc2(); err != nil {    // if문 벗어나면 err 할당 해제

}

// 새로운 에러를 생성해야 하는 경우
return errors.New("Stringlist.ReadFrom: line is too long")

// 새로운 에러를 생성해야 하는 경우
return fmt.Errorf("Stringlist: too long line at %d", count)
```

##  반환하는 변수에 이름을 붙여줄 수도 있음<br>
### return만 사용할 경우, 명명한 결과 인자를 반환함
```Go
// Named return parameter example:
func WriteTo(w io.Writer, lines []string) (n int64, err error) {
    for _, line := range lines {
        var nw int
        nw, err = fmt.Fprintln(w, line)
        n += int64(nw)
        if (err != nil) {

            return  // return with n int 64, err error
        }
    }

    return
}
```

##  함수에 넘겨줄 인자의 갯수가 정해지지 않았을 경우, 가변인자를 사용해 값을 넘겨줄 수 있음 <br>
```Go
// Case #1
func WriteTo(w io.Writer, lines... string) (n int64, err error) {
    ...
}

// Case #2
lines := []string{"Hello", "World", "Go Language"}
WriteTo(w, lines...)
```

##  Go 언어에서 함수는 일급 시민(First-class citizen)으로 분류됨<br>
##  함수는 값으로써 변수에 담길 수 있고, 다른 함수로 넘기거나 돌려받을 수 있음<br>
### 함수 리터럴<br>
```Go
// 함수 리터럴, 익명 함수라고도 불림
func add(a, b int) int {

    return a + b 
}

// 이름이 없는 예:
func(a, b, int) int {

    return a + b
}

// 인자, 반환값이 없는 예:
func printHello() {
    fmt.Println("Hello!")
}

func Example_funcLiteral() {
    func () {
        fmt.Println("Hello!")
    }()
    // Output:
    // Hello!

    return
}

func Example_funcLiteralVar() {
    printHello := func() {
        fmt.Println("Hello!!")
    }
    printHello()
    // Output:
    // Hello!!

    return
}
```

##  고계 함수(higher-order function)는 함수를 넘기고 받는 함수를 나타냄<br>
### 함수를 넘기고 받기 때문에 더 고차원의 함수로 간주된다는 뜻<br>
```Go
package Higher_Order_Func

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {

		return err
	}

	return nil
}

func ExampleReadFrom_Print() {
	r := strings.NewReader("bill\ntom\njane\n")
	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// ( bill )
	// ( tom )
	// ( jane )

	return
}
```
##  클로저(closure)는 외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드를 나타냄<br>
### ReadFrom 함수에서 읽은 line들을 ExampleReadFrom_append 함수에 정의된 문자열 리스트에 추가하는 예제<br>
```Go
func ExampleReadFrom_append() {
    r := strings.NewReader("bill\ntom\njane\n")
    var lines []string
    err := ReadFrom(r, func(line string) {
        lines = append(lines, line)
    })
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(lines)
    // Output:
    // [bill tom jane]
}
}
```
##  생성기(Generator)<br>
### 클로저를 이용한 생성기(generator) 예제 <br>
```Go
func NewIntGenerator() func() int {
    var next int
    return func() int {
        next++
        return next
    }
}

func ExampleNewIntGenerator() {
    gen := NewIntGenerator()
    fmt.Println(gen(), gen(), gen(), gen(), gen())
    fmt.Println(gen(), gen(), gen(), gen(), gen())
    // Output:
    // 1 2 3 4 5
    // 6 7 8 9 10

    return
}

func ExampleNewIntGenerator_multiple() {
    gen1 := NewIntGenerator()
    gen2 := NewIntGenerator()
    fmt.Println(gen1(), gen1(), gen1(), gen1())
    fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
    fmt.Println(gen1(), gen1(), gen1(), gen1())
    // Output:
    // 1 2 3 4 
    // 1 2 3 4 5
    // 5 6 7 8


    return
}
```

## 명명된(named) 자료형<br>
```Go
type rune int32
```

