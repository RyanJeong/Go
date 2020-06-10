package main

import "github.com/ryanjeong/go/cyclic_dependency/after/foo"

func main() {
	o := foo.NewFoo()
	o.PrintFoo()
}
