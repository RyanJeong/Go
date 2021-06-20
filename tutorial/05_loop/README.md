# 반복문

#### Go 언어에서 지원하는 반복문은 `for` 하나 뿐이며, `while`, `do-while`은 사용할 수 없음에 유의할 것

## for 문
### 일반적인 for 문
for (<i>expr1</i>; <i>expr2</i>; <i>expr3</i>) { ... }

* Go 언어에서의 `for`문은 <b>괄호`()`를 사용하지 않음에 유의</b>

```go
package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
```

### <i>expr2</i>만 사용하는 for 문

* <i>expr1</i>, <i>expr3</i>을 생략한 형태

```go
package main

import "fmt"

func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}

```

* <b>표현식 자리에 세미콜론`;` 생략 가능</b>
* 다른 언어에서의 `while`문이 `for`로 표현됨

```go
package main

import "fmt"

func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}

```

### 무한루프 for 문
* `for` 문의 표현식을 모두 생략하면, 무한루프로 동작함
* `for` 문 내부에 별도의 조건을 추가해 무한루프를 벗어나야 함

```go
package main

func main() {
    for {
    }
}
```

### for range 문 (foreach)
* 컬렉션으로부터 요소를 하나씩 가져올 수 있음
* 요소를 하나씩 가져올 때마다 해당 요소의 인덱스와 해당 요소의 값을 가져옴

```go
package main

import "fmt"

func main() {
    colors := []string{"Red", "Green", "Blue"}

    for index, color := range colors {
        fmt.Println(index, color)   // 0 Red
                                    // 1 Green
                                    // 2 Blue
    }
}
```

* 요소로부터 반환되는 값의 순서는 인덱스와 값 순이며, 해당 자리에 `_`를 사용하면 반환되는 값을 무시할 수 있음

```go
package main

import "fmt"

func main() {
    colors := []string{"Red", "Green", "Blue"}

    // You can skip the index or value by assigning to _.
    for _, color := range colors {
        fmt.Println(color)  // Red
                            // Green
                            // Blue
    }
}
```

## break, continue, goto
* `for` 반복분에서 즉시 벗어나야 할 경우 `break` 사용
` `for` 반복문에서 즉시 반복문의 시작 부분(상단 표현식)으로 이동해야 할 경우 `continue` 사용
* 코드를 특정 지역으로 즉시 분기해야 할 경우 `goto` 사용

```go
package main

import "fmt"

func main() {
    a := 0
    for a < 10 {
        if a == 1 {
            a += 4
            continue
        }
        a++
        if a > 5 {
            break
        }
    }
    if a == 6 {
        goto JUMP
    }
    fmt.Println(a) // Can't access this line
JUMP:
    fmt.Println("JUMP") // JUMP
}
```

* `break`와 `label`을 같이 사용할 수 있음 (break label)
* 'break' 문의 `label`은 해당 구조와 <b>무조건 인접해야 함(must be enclosed)</b>
* 아래 경우는 <b>`label`이 `break` 문과 떨어져 있으므로 오류</b>

```go
IGNORE_LOOP: // invalid break label IGNORE_LOOP
    for i := 0; i < 10; i++ {
        for j := 0; j < 10; j++ {
        }
    }

    switch 1 {
    case 1:
        fmt.Println(1)
        for i := 0; i < 10; i++ {
            break IGNORE_LOOP
        }
        fmt.Println(2)
    }
    fmt.Println(3)

```

* `break` 문을 통해 'label'로 분기하면 해당 구조물은 무시됨

```go
package main

import "fmt"

func main() {
IGNORE_LOOP1:
    for i := 0; i < 10; i++ {
        for j := 0; j < 10; j++ {
            fmt.Printf("i=%v, j=%v\n", i, j)
            break IGNORE_LOOP1
        }
    }
    // i=0 j=0

IGNORE_LOOP2:
    switch 1 {
    case 1:
        fmt.Println(1)
        for i := 0; i < 10; i++ {
            break IGNORE_LOOP2
        }
        fmt.Println(2)
    }
    fmt.Println(3)
    // 1
    // 3
}
```
