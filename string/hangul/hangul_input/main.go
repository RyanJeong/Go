package main

import (
	"bufio"
	"fmt"
	"os"
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

func getStr() string {
	in.Scan()

	return in.Text()
}

func main() {
	hangul := getStr()
	index := 0
	for _, r := range hangul {
		index = int(r - start + 1)
	}
	fmt.Fprintln(out, index)
	out.Flush()
}
