```Go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func getInt() int {
	in.Scan()
	s, r := 1, 0
	for _, c := range in.Bytes() {
		if c == '-' {
			s = -1
		} else {
			r = r*10 + int(c-'0')
		}
	}

	return s * r
}

func main() {
	in.Split(bufio.ScanWords)
	n := getInt()
	arr := make([][2]int, n)
	for i := range arr {
		arr[i][0], arr[i][1] = getInt(), getInt()
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
```
