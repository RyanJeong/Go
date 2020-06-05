#   [좌표 정렬하기](https://www.acmicpc.net/problem/11650)<br>
##  문제<br>
### 2차원 평면 위의 점 N개가 주어진다. 좌표를 x좌표가 증가하는 순으로, x좌표가 같으면 y좌표가 증가하는 순서로 정렬한 다음 출력하는 프로그램을 작성하시오.<br>

##  입력<br>
### 첫째 줄에 점의 개수 N (1 ≤ N ≤ 100,000)이 주어진다. 둘째 줄부터 N개의 줄에는 i번점의 위치 x_i와 y_i가 주어진다. (-100,000 ≤ x_i, y_i ≤ 100,000) 좌표는 항상 정수이고, 위치가 같은 두 점은 없다.<br>

##  출력<br>
### 첫째 줄부터 N개의 줄에 점을 정렬한 결과를 출력한다.<br>

##  예제 입력 1<br>
```Text
5
3 4
1 1
1 -1
2 2
3 3
```

##  예제 출력1<br>
```Text
1 -1
1 1
2 2
3 3
3 4
```

##  첫 번째 시도<br>
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

| Memory  | Time    | Bytes      |
| ------- | ------- | ---------- |
| 5512 KB | 2628 ms | 1148 Bytes |

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

| Memory  | Time    | Bytes      |
| ------- | ------- | ---------- |
| 5500 KB | 2432 ms | 646 Bytes  |

##  세 번째 시도<br>
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
```

| Memory  | Time    | Bytes      |
| ------- | ------- | ---------- |
| 5128 KB | 120 ms  | 662 Bytes  |

##  네 번째 시도<br>
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

| Memory  | Time    | Bytes      |
| ------- | ------- | ---------- |
| 4172 KB | 92 ms   | 694 Bytes  |

