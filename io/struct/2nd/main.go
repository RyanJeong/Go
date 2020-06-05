package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Person struct {
	order int
	age   int
	name  string
}

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type ByPersons []Person

func (p ByPersons) Len() int { return len(p) }
func (p ByPersons) Less(i, j int) bool {

	return p[i].age < p[j].age || (p[i].age == p[j].age && p[i].order < p[j].order)
}
func (p ByPersons) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

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
	n := getInt()
	persons := make([]Person, n)
	for i := range persons {
		persons[i] = Person{order: i, age: getInt(), name: getStr()}
	}
	sort.Sort(ByPersons(persons))
	for _, value := range persons {
		fmt.Fprintln(out, value.age, value.name)
	}
	out.Flush()

	return
}
