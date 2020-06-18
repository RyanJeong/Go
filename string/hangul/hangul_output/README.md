```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	start = rune(44032) // 가
	end   = rune(55204) // 힣
)

var in *bufio.Scanner
var out *bufio.Writer

func init() {
	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	in.Split(bufio.ScanWords)
}

func getInt() int {
	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	return n
}

func main() {
	fmt.Fprintln(out, string(start+rune(getInt())-1))
	out.Flush()
}
```
