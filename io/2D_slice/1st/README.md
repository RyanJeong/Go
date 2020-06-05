```Go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type point struct {
	x int
	y int
}

type ByPoint []point

func (a ByPoint) Len() int { return len(a) }
func (a ByPoint) Less(i, j int) bool {
	// if index i's x and index j's x is same:
	if a[i].x == a[j].x {

		return a[i].y < a[j].y
	} else {

		return a[i].x < a[j].x
	}
}
func (a ByPoint) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {
	var n int

	fmt.Fscanf(os.Stdin, "%d", &n)
	points := make([]point, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		points[i].x, _ = strconv.Atoi(string(scanner.Bytes()))
		scanner.Scan()
		points[i].y, _ = strconv.Atoi(string(scanner.Bytes()))
	}
	sort.Sort(ByPoint(points))
	for i := 0; i < n; i++ {
		fmt.Fprintf(os.Stdout, "%d %d\n", points[i].x, points[i].y)
	}

	return
}
```
