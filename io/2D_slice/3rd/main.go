package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	n := 0
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanWords)
	in.Scan()
	n, _ = strconv.Atoi(string(in.Bytes()))
	arr := make([][2]int, n)
	for i := range arr {
		in.Scan()
		arr[i][0], _ = strconv.Atoi(string(in.Bytes()))
		in.Scan()
		arr[i][1], _ = strconv.Atoi(string(in.Bytes()))
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {

			return arr[i][1] < arr[j][1]
		} else {

			return arr[i][0] < arr[j][0]
		}
	})
	for _, value := range arr {
		fmt.Fprintln(out, value[0], value[1])
	}
	out.Flush()

	return
}
