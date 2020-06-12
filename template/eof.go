/* Ctrl + Shift + D */
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

func main() {
	count := 0
	for in.Scan() {
		n, _ := strconv.Atoi(in.Text())
		if n > 0 {
			count++
		}
	}
	fmt.Fprintln(out, count)
	out.Flush()

	return
}
