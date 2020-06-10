package foo

import (
	"fmt"

	"github.com/ryanjeong/go/cyclic_dependency/before/bar"
)

type Foo struct {
}

func (foo Foo) PrintFoo() {
	fmt.Println(foo)
}

func NewFoo() *Foo {
	foo := new(Foo)
	return foo
}

func RequireBar() {
	o := bar.NewBar()
	o.PrintBar()
}
