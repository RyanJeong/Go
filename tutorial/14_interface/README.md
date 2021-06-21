# [WIP] 인터페이스
* 구조체(struct)가 필드들의 집합체라면, interface는 메서드들의 집합체다. 
* <i>interface type</i>은 메소드 시그니쳐(method signature, 메소드의 이름과 매개변수만 표현한 형태)들의 집합 
* 구조체와 마찬가지로, type 키워드를 사용해 인터페이스를 정의한다.
* interface type 값은 구현된 메소드들을 값으로써 저장할 수 있음 

* <b>struct type에서의 메소드는, pointer receiver를 써서 선언할 경우, pointer receiver와 value receiver에서 모두 작동한다. 그러나 인터페이스는 다르다.

인터페이스의 메소드 자체를 호출할때, 런타임에서 동적타입이 고려되고, 동적 값이 메소드의 리시버로 넘겨지게 된다. s의 동적타입은 Rect 이고, Rect는 Area 메소드를 구현하는것이 아닌, *Area를 구현한다.</b>
* 추가 공부 : [링크](https://medium.com/rungo/interfaces-in-go-ab1601159b3a)

```go
package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    a = &v // a *Vertex implements Abser

    // In the following line, v is a Vertex (not *Vertex)
    // and does NOT implement Abser.
    a = v

    fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

* Interface가 가진 메서드를 갖고있는 타입들은 Interface의 구성원인 것처럼 처리된다. 상속과 비슷하게 동작하는데, 위의 예제에서는 MyFloat와 *Vertex가 Abs()라는 메서드를 가지고 있기 때문에 Abser의 구성원이 될 수 있다는 것이다.
* 'Interfaces are implemented implicitly(인터페이스는 암시적으로 충족 가능하다)'
* 사용자 정의 타입이 인터페이스에 관련되어있다고 명시하지 않더라도 연관된 메서드가 있으면 인터페이스가 해당 구현 메서드를 사용할 수 있는 상태가 된다. 

```go
package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
    fmt.Println(t.S)
}

func main() {
    var i I = T{"hello"}
    i.M()
}

```

### Type Assertion
* 빈 인터페이스는 C언어에서의 `void *`와 동일

```go
package main
 
import "fmt"
 
func main() {
    var x interface{}
    x = 1 
    x = "Tom"
 
    printIt(x)
}
 
func printIt(v interface{}) {
    fmt.Println(v) //Tom
}
```

* A <i>type assertion</i> provides <b>access to an interface value's underlying concrete value</b>.
* 실제 값에 접근할 수 있게 해줌 

```go
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println(s) // hello

    s, ok := i.(string)
    fmt.Println(s, ok) // hello true

    f, ok := i.(float64)
    fmt.Println(f, ok) // 0 false

    f = i.(float64) // panic: interface conversion: interface {} is string, not float64
    fmt.Println(f)
}

```

```go
package main

import "fmt"

func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

func main() {
    do(21)
    do("hello")
    do(true)
}
```

## Interface values
```go
package main

import (
    "fmt"
    "math"
)

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    fmt.Println(t.S)
}

type F float64

func (f F) M() {
    fmt.Println(f)
}

func main() {
    var i I

    i = &T{"Hello"}
    describe(i)
    i.M()

    i = F(math.Pi)
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}

```

### Interface values with nil underlying values

```go
package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

func main() {
    var i I

    var t *T
    i = t
    describe(i)
    i.M()

    i = &T{"hello"}
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
```

* <b>A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.</b>

### interface{}

* An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

```go
package main

import "fmt"

func main() {
    var i interface{}
    describe(i)

    i = 42
    describe(i)

    i = "hello"
    describe(i)
}

func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}

```
### Stringers

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

```go
type Stringer interface {
    String() string
}
```

A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
    a := Person{"Arthur Dent", 42}
    z := Person{"Zaphod Beeblebrox", 9001}
    fmt.Println(a, z)
}
```
