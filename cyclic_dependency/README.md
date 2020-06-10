# Cyclic Dependency in Golang
```Text
// Before
//
// Foo depends on Bar
// Bar depends on Foo
import cycle not allowed
package main
	imports github.com/ryanjeong/go/cyclic_dependency/before/foo
	imports github.com/ryanjeong/go/cyclic_dependency/before/bar
	imports github.com/ryanjeong/go/cyclic_dependency/before/foo
```
```Text
// After
// The cycle is broken;
{}
```

