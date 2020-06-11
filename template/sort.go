package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	return n
}

func getStr() string {
	in.Scan()

	return in.Text()
}

func main() {
	var start, end, mid int
	x, y := getInt(), getInt()
	arr1, arr2 := make([]string, x), make([]string, 0, y)
	for i := range arr1 {
		arr1[i] = getStr()
	}
	/* sort.Strings(arr1) */
	sort.Slice(arr1, func(i, j int) bool {

		return arr1[i] < arr1[j]
	})
	for i := 0; i < y; i++ {
		str := getStr()
		start = 0
		end = x - 1
		for start <= end {
			mid = (start + end) / 2
			if arr1[mid] == str {
				arr2 = append(arr2, str)
				break
			} else if arr1[mid] > str {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	/* sort.Strings(arr2) */
	sort.Slice(arr2, func(i, j int) bool {

		return arr2[i] < arr2[j]
	})
	fmt.Fprintln(out, len(arr2))
	for i := range arr2 {
		fmt.Fprintln(out, arr2[i])
	}
	out.Flush()

	return
}
