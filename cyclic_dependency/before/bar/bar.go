package bar

import (
	"fmt"

	"github.com/ryanjeong/go/cyclic_dependency/before/foo"
)

type Bar struct {
}

func (bar Bar) PrintBar() {
	fmt.Println(bar)
}

func NewBar() *Bar {
	bar := new(Bar)

	return bar
}

func RequireFoo() {
	o := foo.NewFoo()
	o.PrintFoo()
}
