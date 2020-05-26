package convert

import (
	"fmt"
	"strconv"
)

func Example_strconv() {
	var i int
	var l int64
	var f float64
	var s string
	var err error

	i, err = strconv.Atoi("350")
	if err == nil {
		fmt.Println(i) // 350
	}
	l, err = strconv.ParseInt("cc7fdd", 16, 32) // strconv.ParseInt(s, base, bitSize)
	if err == nil {
		fmt.Println(l) // 13402077
	}
	// If base == 0, the base is implied by the string's prefix:
	//     base 16 for "0x", base 8 for "0", and base 10 otherwise.
	l, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	if err == nil {
		fmt.Println(l) // 13402077
	}
	f, err = strconv.ParseFloat("3.14", 64)
	if err == nil {
		fmt.Println(f) // 3.14
	}
	s = strconv.Itoa(340)
	if err == nil {
		fmt.Println(s) // "340"
	}
	s = strconv.FormatInt(13402077, 16)
	if err == nil {
		fmt.Println(s) // "cc7fdd"
	}
	// Output:
	// 350
	// 13402077
	// 13402077
	// 3.14
	// 340
	// cc7fdd
}

func Example_fmt() {
	var num int
	var s string

	fmt.Sscanf("57", "%d", &num)
	fmt.Println(num)     // 57
	s = fmt.Sprint(3.14) // "3.14"
	fmt.Println(s)
	s = fmt.Sprintf("%x", 13402077) // "cc7fdd"
	fmt.Println(s)
	// Output:
	// 57
	// 3.14
	// cc7fdd
}
