```Text
# May 26 2020, Working directory:
./
├── cmd
│   └── fib
│       ├── fib.go
│       └── go.sh
├── go.sh
└── seq.go
```
```Go
// Package seq implemets functions for well-known sequences like Fibonacci.
package seq

// Fib returns n-th (from 0-th) Fibonacci number.
func Fib(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		p, q = q, (p + q)
	}

	return p
}
```
