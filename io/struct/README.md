#   [나이순 정렬](https://www.acmicpc.net/problem/10814)<br>
##  문제<br>
### 온라인 저지에 가입한 사람들의 나이와 이름이 가입한 순서대로 주어진다. 이때, 회원들을 나이가 증가하는 순으로, 나이가 같으면 먼저 가입한 사람이 앞에 오는 순서로 정렬하는 프로그램을 작성하시오.<br>

##  입력<br>
### 첫째 줄에 온라인 저지 회원의 수 N이 주어진다. (1 ≤ N ≤ 100,000) 둘째 줄부터 N개의 줄에는 각 회원의 나이와 이름이 공백으로 구분되어 주어진다. 나이는 1보다 크거나 같으며, 200보다 작거나 같은 정수이고, 이름은 알파벳 대소문자로 이루어져 있고, 길이가 100보다 작거나 같은 문자열이다. 입력은 가입한 순서로 주어진다.<br>

##  출력<br>
### 첫째 줄부터 총 N개의 줄에 걸쳐 온라인 저지 회원을 나이 순, 나이가 같으면 가입한 순으로 한 줄에 한 명씩 나이와 이름을 공백으로 구분해 출력한다.<br>

##  예제 입력 1<br>
```Text
3
21 Junkyu
21 Dohyun
20 Sunyoung
```

##  예제 출력1<br>
```Text
20 Sunyoung
21 Junkyu
21 Dohyun
```

##  첫 번째 시도<br>
```Go
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
```

| Memory  | Time    | Bytes      |
| ------- | ------- | ---------- |
| 6460 KB | 88 ms   | 852 Bytes  |

##  두 번째 시도<br>
```Go
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
```

| Memory  | Time    | Bytes      |
| ------- | ------- | ---------- |
| 6988 KB | 80 ms   | 1241 Bytes |
