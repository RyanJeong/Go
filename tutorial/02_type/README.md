# 타입 (data types)

## Go 언어에서는 다음과 같이 기본적인 타입 제공
### boolean
* `bool`

### string
* `string`
* `string` 타입으로 선언된 변수는 일종의 구조체이며, 멤버 변수로 실제 문자열을 가리키는 포인터`ptr`와 문자열의 길이를 저장하는 변수`len`으로 구성되어 있음
* string 타입의 `ptr`이 가리키는 대상은 <b>수정 불가(immutable)</b>

### integer
* `int`
* `int8`
* `int16`
* `int32`
* `int64`
* `uint`
* `uint8`
* `uint16`
* `uint32`
* `uint64`
* `uintptr`

#### `int`, `uint`, `uintptr` 타입은 32비트 환경에서는 32비트, 64비트 환경에서는 64비트가 적용됨
#### 특별한 이유가 없는 한, 정수 값 표현에는 `int` 타입을 사용해야 하며, 정수의 타입 추론 결과는 `int` 타입임

```go
i := 32 // int
```

### floating point 또는 complex number
* 'float32'
* 'float64'
* 'complex64'
* 'complex128'

#### 실수값의 타입 추론 결과는 `float64`, 복소수값의 타입 추론 결과는 `complex128`임 

```go
f := 3.142 // float64
g := 0.867+0.5i // complex128
```

### 기타
* `byte` (alias for uint8)

#### `byte` 타입은 ASCII 코드를 표현할 때 사용

```go
a := 'A'; // A's ASCII code is 0x41
```
<br>

* `rune` (alias for int32)

#### `rune` 타입은 유니코드 코드 포인트(Unicode code point)를 표현하기 위해 사용

```go
r := '홍' // '홍's Unicode code point is U+D64D
```

## 문자와 문자열
### Single quote `''`
* 문자 하나 표현 

```go
	a := 'a'
	fmt.Printf("%c", a) // a
```

### Double quote `""`
* 문자열 표현

```go
	str := "Hello\n"
	fmt.Println(str) // Hello
                     // (newline)
```

### Back quote ``````
* 문자열을 표현하되, 이스케이프 시퀀스를 무시함

```go
	str_raw := `Hello\nWorld`
	fmt.Println(str_raw) // Hello\nWorld
```

#### 문자열 정의 시 여러 줄을 사용해야 하는 경우 사용됨
#### `""`와 `+`를 같이 사용해 여러 줄을 사용할 수도 있음

```go
    line1 := `Hello
World`
    line2 := "Hello\n" + "World"
    fmt.Println(line1)  // Hello
                        // World

    fmt.Println(line2)  // Hello
                        // World
```

## 타입 변환 (type conversion)
* 표현식 `T(v)`는 값 `v`를 `T` 타입으로 변환

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

* 더욱 간단히 표현하면 아래와 같음

```go
i := 42
f := float64(i)
u := uint(f)
```
