package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var in *bufio.Scanner
var out *bufio.Writer

func init() {
	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	in.Split(bufio.ScanWords)
}

func getInt() int64 {
	in.Scan()
	n, _ := strconv.ParseInt(in.Text(), 10, 64)

	return n
}

func getStr() string {
	in.Scan()

	return in.Text()
}

func main() {
	fmt.Fprintln(out, getInt()*4)
	out.Flush()
}
