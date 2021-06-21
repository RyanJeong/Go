# 클로져 (closure)
* 특정 함수 내부에 있는 익명함수가 해당 함수의 지역변수에 참조할 수 있는 경우, 해당 익명함수를 클로져라 함
* in this sense the function is "bound" to the variables.
* C 언어에서의 `static` 지역변수와 흡사하지만, 함수가 값으로써 서로 다른 변수에 할당됨에 따라 각 변수가 가리키는 함수는 서로 다른 문맥(context)를 사용하게 됨

```go
package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
            /*
            0 0
            1 -2
            3 -6
            6 -12
            10 -20
            15 -30
            21 -42
            28 -56
            36 -72
            45 -90
            */
        )
    }
}
```
#### 해설
##### 1. `pos`, `neg`는 `adder` 함수의 반환형을 할당받음
##### 2. `adder` 함수의 반환형은 익명함수 `func(x int) int { ... }`
##### 3. 해당 익명함수는 `adder` 함수의 지역변수 `sum`을 사용
##### 4. 익명함수는 클로져로 처리되며, 해당 익명함수는 `added` 함수 내부에 inline 형태로 처리됨과 동시에 변수 `sum`을 stack이 아닌 heap에 할당
###### 4-1. `pos`, `neg`는 `adder` 함수의 지역변수인 `sum`에 접근이 가능해짐
###### 4-2. `pos`, `neg`는 서로 다른 메모리에 할당되며, `pos`가 접근하는 변수 `sum`과 `neg`가 접근하는 변수 `sum`는 서로 독립적임
##### 5. `pos`와 `neg`에 전달인자를 넘겨주면, 클로저인 익명함수의 매개변수로 전달되고, 전달된 매개변수는 해당 문맥에 맞는 변수 `sum`에 접근해 값을 처리한 후 해당 값을 반환하게 됨
