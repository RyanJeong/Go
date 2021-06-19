# 조건문
## if 문
* `if` 문에서 조건을 나타내는 표현식은 <b>boolean 타입을 반환</b>해야 하며, 해당 표현식의 양 옆에 괄호`()`가 필요하지 않음
* `if` 문 블록의 시작을 나타내는 중괄호`{`는 `if 문과 같은 줄에 표현되어야 함

```go
    x := 1
    if x==1 {
		fmt.Println("vim-go")
    }
```

* `else` 또는 `else if` 문은 블록의 종료를 나타내는 중괄호'}' 다음에 등장해야 함
* `else` 또는 `else if` 문 블록의 시작을 나타내는 중괄호'{'는 `else` 또는 `else if` 문과 같은 줄에 표현되어야 함 (cuddled else)

```go
    x := 3
    if x==1 {
		fmt.Println("vim-go")
    } else if x==2 {
		fmt.Println("vim")
    } else {
		fmt.Println("go")
    }
```

* `if` 문의 표현식 자리에 간단한 문장(optional statement)를 사용할 수 있음
* 해당 문장 안에 변수를 선언할 수 있으며, 이 변수의 유효 범위는 `if` 문 이내임
* `if` 문에서 선언한 변수는 `if` 문을 벗어나면 사용할 수 없음 (블록 종료와 동시에 반환됨)

```go
	i := 4
	if in := i * i; in == 16 {
		fmt.Println("go")
	}
```

## [WIP]switch 문
* switch 문의 조건을 나타내는 표현식과 일치하는 곳으로 코드가 분기됨
* Go 언어의 switch 문은 각 case의 마지막에 자동으로 `break` 문장이 추가됨
* switch 문의 

```go
    var msg string
    state := 2
    
    switch state {
    case 1:
        msg = "Good" 
    case 2:
        msg = "So-so"
    case 3, 4: // case를 콤마로 구분해 중첩할 수 있음 
        msg = "Bad!"
    default:
        msg = "Invalid"
    }
    fmt.Println(msg) // So-so
```

* if에서와 마찬가지로, 표현식 자리에 간단한 문장이 올 수 있음

```go
	switch new_state := state*2 - 5; new_state { // -1
	case 1:
		msg = "Good"
	case 2:
		msg = "So-so"
	case 3, 4:
		msg = "Bad!"
	default:
		msg = "Invalid"
	}
	fmt.Println(msg) // Invalid
```

* switch 문에 조건이 생략될 수 있음
* 조건문이 생략된 switch 문은 매우 긴 else-if 구문을 대치할 수 있음
* 조건문이 생략되면, switch에서의 조건은 `true`이며, 각 case의 표현식이 참인지 아닌지를 판별해 표현식이 참인 곳으로 코드가 분기됨

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}
```

