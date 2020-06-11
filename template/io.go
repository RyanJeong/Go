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

func getInt() int {
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	/*
		n := 0
		for _, c := range in.Bytes() {
			n = (n * 10) + int(c-'0')
		}
	*/

	return n
}

func getStr() string {
	in.Scan()

	return in.Text()
}

func main() {
	cloud := 0
	x, y := getInt(), getInt()
	arr := make([]int, x*y)
	for i := range arr {
		arr[i] = -1
	}
	for i := 0; i < x; i++ {
		str := getStr()
		cloud = 0
		for j, c := range str {
			if string(c) == "c" {
				cloud++
				arr[i*y+j] = 0
			} else if cloud != 0 {
				arr[i*y+j] = arr[i*y+j-1] + 1
			}
		}
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Fprintf(out, "%d ", arr[i*y+j])
		}
		fmt.Fprintln(out)
	}
	out.Flush()

	return
}
