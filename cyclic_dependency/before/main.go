package main

import "github.com/ryanjeong/go/cyclic_dependency/before/foo"

func main() {
	o := foo.NewFoo()
	o.PrintFoo()
}
