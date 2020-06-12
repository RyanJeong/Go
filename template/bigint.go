package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var in *bufio.Scanner
var out *bufio.Writer

func init() {
	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	in.Split(bufio.ScanWords)
}

func main() {
	x, y, sum := new(big.Int), new(big.Int), new(big.Int)
	in.Scan()
	x, xOk := x.SetString(in.Text(), 10)
	if !xOk {
		fmt.Println("SetString: Error")
	}
	in.Scan()
	y, yOk := y.SetString(in.Text(), 10)
	if !yOk {
		fmt.Println("SetString: Error")
	}
	fmt.Fprintln(out, sum.Add(x, y))
	out.Flush()

	return
}
