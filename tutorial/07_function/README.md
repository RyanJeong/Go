# 함수와 익명함수
## 함수 정의
* `func` 키워드 사용
* 함수 매개변수는 0개 이상 나열할 수 있음
* <b>반환형은 매개변수 다음에 위치함에 유의</b>

```go
package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
```

* 함수의 매개변수는 다음과 같이 축약할 수 있으며, <b>목록 마지막에는 자료형을 꼭 기입해야 함</b>

```go
package main

import "fmt"

func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
```

* Go 언어에서의 함수는 값을 여러 개 반환할 수 있음

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}

```

* Go 언어에서의 함수는 값을 반환할 때 지역변수의 이름으로 반환할 수 있음
* `return` 다음에 표현식이 생략되면, 해당 함수의 지역변수를 사용해 반환할 수 있으며, 이러한 반환 형태를 "naked" return이라고 표현함
* 반환형 자리에 지역변수 이름을 같이 표기하며, 해당 이름에는 의미있는 값이 기록되야 함
* 매개변수 목록의 축약 표현이 가능한 것처럼, 반환할 변수들도 아래와 같이 축약 표현이 가능함 
* <b>이러한 형태의 코드는 가독성을 해치므로, 몇 줄 안 되는 간결한 함수에서만 사용할 것을 권장</b>

```go
package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
```

### Call by value
* 값을 복사해 넘겨줌
* 전달인자로 넘겨준 변수는 호출한 함수의 매개변수로 복사되며, 매개변수는 전달인자와 서로 다른 메모리를 사용하는 독립적인 변수임 


### Call by address
* 값의 주소를 넘겨줌
* 전달인자로 넘겨준 변수는 호출한 함수의 매개변수로 복사되며, 매개변수는 전달인자와 서로 다른 메모리를 사용하는 독립적인 변수이면서 해당 전달인자에 간접 참조할 수 있는 상태가 됨

```go
package main

import "fmt"

func main() {
    i := 1
    fmt.Println("Before", i) // 1
    val(i)
    fmt.Println("After ", i) // 1
    fmt.Println("Before", i) // 1
    ref(&i)
    fmt.Println("After ", i) // 2
}

func val(i int) {
    i++
    return
}

func ref(i *int) {
    *i++
    return
}
```

### 가변인자 함수 (variadic function)
* 함수의 매개변수를 고정하지 않고 사용
* 문자열 가변 매개변수를 사용하는 예:

```go
package main

import "fmt"

func main() {
    say("Hello", "World", "Golang") // Hello
                                    // World
                                    // Golang
    say("Variadic Function") // Variadic Function 
}

func say(msg ...string) {
    for _, s := range msg {
        println(s)
    }
}

```

* 일반적인 매개변수와 가변 길이의 매개변수를 같이 사용하는 예:

```go
package main

import "fmt"

func main() {
    fmt.Println(calc("Rectangle", 10, 20)) // 200
    fmt.Println(calc("Square", 5))         // 25
}


func calc(str string, y ...int) int {
    area := 1

    for _, val := range y {
        if str == "Rectangle" {
            area *= val
        } else if str == "Square" {
            area = val * val
        }
    }

    return area
}
```

## 익명함수 (anonymous function)
* 이름이 생략된 함수
* Go 언어에서는 함수를 값처럼 취급하며, 따라서 변수에 기록될 수 있음
    * Go 언어에서의 함수는 일급함수로 Go 언어에서 기본으로 제공되는 타입과 동일하게 취급됨
    * 변수에 할당, 매개변수로의 전달, 반환값 등으로 사용할 수 있음
* 익명함수는 변수에 할당되거나 다른 함수의 매개변수에 전달될 수 있음
* 일반 변수에 익명함수를 할당한 예:

```go
package main

import "fmt"

func main() {
    sum := func(n ...int) int {
        s := 0
        for _, i := range n {
            s += i
        }
        return s
    }
    result := sum(1, 2, 4, 8, 16, 32) // 63
    fmt.Println(result)
}
```

* 익명함수를 다른 함수의 매개변수에 전달하고자 할 경우, 익명함수의 형태와 익명함수를 전달받을 매개변수의 형태가 일치해야 함
    * <b>익명함수의 매개변수 목록</b>과 <b>익명함수를 전달받을 매개변수의 매개변수 목록</b>이 일치해야 함
    * 익명함수의 반환형과 익명함수를 전달받을 매개변수의 반환형이 일치해야 함
    
* 다른 함수의 매개변수에 익명함수를 전달한 예:

```go
package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))

    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
}
```

## type 키워드
* `type`은 C 언어에서의 `typedef`와 유사
* UDT(user defined type)을 정의하기 위해 사용
* `type`을 사용해 `func(int, int) int`를 `calculator`로 별명을 붙인 예:

```go
package main

type calculator func(int, int) int

func main() {
    add := func(i int, j int) int {
        return i + j
    }

    r1 := calc(add, 10, 20)
    println(r1) // 30

    r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
    println(r2) // -10
}

// func calc(f func(int, int) int, a int, b int) int {
func calc(f calculator, a int, b int) int {
    result := f(a, b)
    return result
}
```
