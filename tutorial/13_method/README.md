# [WIP] 메서드

* Remember: a method is just a function woth a <i>receiver</i> argument.
* The receiver appears in it own argument list between the `func` keyword and the method name.

```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
}

```

* Here's `Abs` written as a regular function with no change in functionality.

```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(Abs(v))
}
```

* You can declare a method on non-struct types, too.<br>
In this example we see a numeric type `MyFloat` with an `Abs` method.<br>
You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as `int`).

```go
package main

import (
    "fmt"
    "math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}

```

* Go 언어는 객체지향 프로그래밍(OOP)을 고유의 방식으로 지원한다. 타 언어의 OOP의 클래스가 필드와 메서드를 함께 갖는 것과 달리 Go 언어에서는 struct가 필드만을 가지며, 메서드는 별도로 분리되어 정의된다.

* Go 메서드는 특별한 형태의 func 함수이다. 메서드는 함수 정의에서 func 키워드와 함수명 사이에 "그 함수가 어떤 struct를 위한 메서드인지"를 표시하게 된다. 흔히 receiver로 불리우는 이 부분은 메서드가 속한 struct 타입과 struct 변수명을 지정하는데, <b>struct 변수명은 함수 내에서 마치 입력 파라미터처럼 사용된다</b>.
* 예를 들어, 아래 예제는 Rect라는 struct를 정의하고 area() 라는 메서드를 정의하고 있다. func와 area() 사이에 Rect 타입의 r 이 정의되고 이를 함수 본문에서 사용하고 있다. 메서드가 선언된 이후에는 Rect 구조체의 객체는 rect.area() 문장처럼 area() 메소드를 struct 객체로부터 직접 호출할 수 있다.

```go
package main
 
//Rect - struct 정의
type Rect struct {
    width, height int
}
 
//Rect의 area() 메소드
func (r Rect) area() int {
    return r.width * r.height   
}
 
func main() {
    rect := Rect{10, 20}
    area := rect.area() //메서드 호출
    println(area)
}
```

* Value receiver는 struct의 데이타를 복사(copy)하여 전달하며, 포인터 receiver는 struct의 포인터만을 전달한다. Value receiver의 경우 만약 메서드내에서 그 struct의 필드값이 변경되더라도 호출자의 데이타는 변경되지 않는 반면, 포인터 receiver는 메서드 내의 필드값 변경이 그대로 호출자에서 반영된다. (call by value vs. call by address)

* 위의 Rect.area() 메서드는 Value receiver를 표현한 것으로 만약 area() 메서드 내에서 width나 height가 변경되더라도 main() 함수의 rect 구조체의 필드값에는 변화가 없다. 하지만, 아래와 같이 이를 포인터 receiver로 변경한다면, 메서드 내 r.width++ 필드 변경분이 main() 함수에서도 반영되기 때문에 출력값이 11, 220 이 된다.

```go
// 포인터 Receiver
func (r *Rect) area2() int {
    r.width++
    return r.width * r.height
}
 
func main() {
    rect := Rect{10, 20}
    area := rect.area2() //메서드 호출
    println(rect.width, area) // 11 220 출력
}
```
* The `Scale` method operates on a copy of the original `Vertex` value. 
* The `Scale` method <b>must have a pointer receiver to change the Vertex value</b> declared in the main function.
```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(10)
    fmt.Println(v.Abs())
}
```

## method vs function in pointer

```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    Scale(&v, 10) // need a operator &
    fmt.Println(Abs(v))
}

```

* For the statement `v.Scale(5)`, even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the Scale method has a pointer receiver.
```go
package main

import "fmt"

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.Scale(2)        // not need a operator &
    // v is a value and not a pointer, the method with the pointer receiver is called automatically
    ScaleFunc(&v, 10) // need a operator &

    p := &Vertex{4, 3}
    p.Scale(3)
    ScaleFunc(p, 8)   // not need a operator

    fmt.Println(v, p)
}

```

* in reverse direction: 

```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs())
    fmt.Println(AbsFunc(v))

    p := &Vertex{4, 3}
    fmt.Println(p.Abs())
    fmt.Println(AbsFunc(*p))
}
```

* Functions that take a value argument must take a value of that specific type: 

```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

* while methods with value receivers take either a value or a pointer as the receiver when they are called:

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

* In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.

## Choosing a value of pointer receiver
* There are two reasons to use a pointer receiver.

* The first is so that the method can modify the value that its receiver points to.

* The second is <b>to avoid copying the value on each method call</b>. This can <b>be more efficient</b> if the receiver is a large struct, for example.

```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := &Vertex{3, 4}
    fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
    v.Scale(5)
    fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

* In this example, both `Scale` and `Abs` are with receiver type `*Vertex`, <b>even though the `Abs` method needn't modify its receiver</b>.

* In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)

# 아래 내용 정리!
* [Golang is no pass-by-reference in Go](https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go)
