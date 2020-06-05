package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type person struct {
	order int
	age   int
	name  string
}

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

func getStr() string {
	in.Scan()

	return string(in.Bytes())
}

func main() {
	in.Split(bufio.ScanWords)
	n := getInt()
	arr := make([]person, n)
	for i := range arr {
		arr[i].order, arr[i].age, arr[i].name = i, getInt(), getStr()
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].age == arr[j].age {

			return arr[i].order < arr[j].order
		} else {

			return arr[i].age < arr[j].age
		}
	})
	for _, value := range arr {
		fmt.Fprintln(out, value.age, value.name)
	}
	out.Flush()

	return
}
