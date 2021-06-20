# 포인터

* Go 언어는 포인터를 지원하며, 포인터 변수는 변수의 메모리 상 주소 값을 보관
* `*T` 형은 `T` 형을 가리키는 포인터
* 포인터에서의 초기치는 `nil`

```go
    var p *int // init. to <nil>
```

* `&` 연산자는 해당 변수를 가리킬 수 있는 주소 반환

```go
    i := 42
    p = &i
```

* `*` 연산자는 포인터가 가리키는 대상의 실제 값 참조 (<i>dereferencing</i> or <i>indirecting</i>)

```go
    fmt.Println(*p) // read i through the pointer p 
    *p = 21         // set i through the pointer p
```

* <b>Go 언어는 C 언어처럼 포인터 주소 연산을 지원하지 않음에 유의</b>

```go
    fmt.Println(p + 1) // invalid operation: p + 1
    fmt.Println(p++)   // syntax error, ++ is not an operator
```

```go
package main

import "fmt"

func main() {
    var p *int
    fmt.Println(p) // <nil>

    i := 42 
    p = &i // p points to i

    fmt.Println(*p) // 42, read i through the pointer
    fmt.Println(p)  // 0xc0000a4018, print i's address
    *p = 21         // set i through the pointer
    fmt.Println(*p) // 21, read the new value of i
    fmt.Println(p)  // 0xc0000a4018
}
```

