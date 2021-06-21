package main

import "fmt"

// struct 정의
type person struct {
	name string
	age  int
}

type Vertex struct {
	X int
	Y int
}

type Vertex1 struct {
	X, Y int
}

var (
	v1 = Vertex1{1, 2}  // has type Vertex1
	v2 = Vertex1{X: 1}  // Y:0 is implicit
	v3 = Vertex1{}      // X:0 and Y:0
	p  = &Vertex1{1, 2} // has type *Vertex1
)

type dict struct {
	data map[int]string
}

//생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d //포인터 전달
}

func main() {
	// person 객체 생성
	p := person{}

	// 필드값 설정
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 5
	fmt.Println(v.X)

	p_v := &v
	p_v.X = 1e9 // don't use ->, structure's dereferencing is used to '.'
	fmt.Println(p_v)

	fmt.Println(v1, v2, v3, p)

	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}

	p3 := new(person)
	p3.name = "Lee" // p가 포인터라도 . 을 사용한다
	fmt.Println(p1, p2, p3)

	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
}
