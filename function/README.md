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
```Go
func WriteTo(w io.Writer, lines []string) (int64, error)

// 함수 호출 후 반환되는 예외값을 바로 처리할 수 있는 패턴
if err := MyFunc(); err != nil {    // if문 벗어나면 err 할당 해제

}

if _, err := MyFunc2(); err != nil {    // if문 벗어나면 err 할당 해제

}
```

##  반환하는 변수에 이름을 붙여줄 수도 있음<br>
```Go
// Named return parameter example:
func WriteTo(w io.Writer, lines []string) (n int64, err error) {
    ...
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

