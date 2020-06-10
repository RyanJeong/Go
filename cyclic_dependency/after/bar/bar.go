package bar

import (
	"fmt"
)

type I_foo interface {
	PrintFoo()
}

type Bar struct {
}

func (bar Bar) PrintBar() {
	fmt.Println(bar)
}

func NewBar() *Bar {
	bar := new(Bar)

	return bar
}

func RequireFoo(o I_foo) {
	o.PrintFoo()
}
