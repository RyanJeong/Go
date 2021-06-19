# 변수와 상수

## 변수
* Go 언어에서 변수를 선언하고자 할 때, `var` 사용
* `var`, 변수 명, 타입 순으로 작성

```go
var a int
```

* 변수 선언과 동시에 초기화하고자 할 경우 `=` 사용

```go
var f float32 = 11.0
```

* 선언된 변수는 해당 변수의 타입과 일치하는 값을 대입을 통해 저장될 수 있음

```go
a = 10 // variable a is int type
f = 12.0 // variable f is float64 type
```

* 동일한 타입 변수를 여러 개 선언하고자 할 경우 `var`과 타입 사이에 목록을 나열함

```go
var a, b, c int
```

* 여러 개 선언한 변수들을 선언함과 동시에 초기화하고자 할 경우, `=` 사용

```go
var a, b, c int = 1, 2, 3
```

* Go 언어는 타입 추론이 가능하므로 아래와 같이 사용할 수 있음

```go
var i = 1 // variable i is int type
var s = "Go" // variable s is string type
```

* Go 언어는 short assignment statement `:=`를 지원하며, 변수 선언을 아래와 같이 간결히 표현할 수 있음
    * 타입 추론이 가능하다는 특징을 확장해 간결한 연산을 지원
    * short assignment statement는 <b>함수 내에서만 사용 가능</b>하며, 함수 밖에서 변수를 선언하고자 할 경우 `var`을 꼭 사용해야 함

```go
i := 1
```

* Go 언어에서는 변수를 선언하고 한 번도 사용하지 않으면 <b>오류</b>가 발생함

* 변수 선언 시 초기치가 생략되어 있다면 자동으로 초기화 수행
    * 만약 정수타입 변수라면 해당 변수는 0으로, `bool` 타입이라면 false, `string` 타입이라면 ""(empty string)으로 초기화

* Go 언어의 변수와 상수는 함수 밖에서도 사용 가능

## 상수
* Go 언어에서 상수를 선언하고자 할 때, `const` 사용
* `const`, 변수 명, 타입 순으로 작성

```go
const MAX int = 10
const MSG string = "HI"
```

* Go 언어는 타입 추론이 가능하므로 아래와 같이 사용할 수 있음

```go
const MAX = 10
const MSG = "HI"
```

* 상수를 여러 개 선언하고자 할 경우, `const`를 묶어서 사용할 수 있음

```go
const (
    MAX = 10
    MIN = 0
    MSG = "HELLO"
)
```

* Go 언어에서 상수를 선언할 때, `iota`를 사용하면 값을 0부터 순차적으로 초기화할 수 있음

```go
const (
    LOW = iota  // 0
    MEDIUM      // 1
    HIGH        // 2
)
```

* `iota` 응용 예

```go
type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```

## 키워드(keyword)

* 아래 키워드들은 변수명, 상수명, 함수명 등으로 사용할 수 없음(identifier로 사용 불가)

```text
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```
