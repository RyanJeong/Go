```Go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int

	fmt.Fscanf(os.Stdin, "%d", &n)
	arr := make([][2]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i][0], _ = strconv.Atoi(string(scanner.Bytes()))
		scanner.Scan()
		arr[i][1], _ = strconv.Atoi(string(scanner.Bytes()))
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {

			return arr[i][1] < arr[j][1]
		} else {

			return arr[i][0] < arr[j][0]
		}
	})
	for i := 0; i < n; i++ {
		fmt.Fprintf(os.Stdout, "%d %d\n", arr[i][0], arr[i][1])
	}

	return
}
```
