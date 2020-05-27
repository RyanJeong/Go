```Text
$ go test -bench=.
```
```Text
goos: linux
goarch: amd64
pkg: github.com/ryanjeong/go/string/concat/benchmark
BenchmarkSprintf4-12    	10000000	       215 ns/op
BenchmarkPlus4-12       	50000000	        23.8 ns/op
BenchmarkSprint4-12     	10000000	       222 ns/op
BenchmarkJoin4-12       	20000000	        99.5 ns/op
BenchmarkBytes-12       	20000000	       100 ns/op
PASS
ok
```
