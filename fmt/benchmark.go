package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	w := bufio.NewWriter(os.Stdout)

	s1 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Printf("P")
	}
	t1 := time.Since(s1)

	s2 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Fprintf(w, "F")
	}
	t2 := time.Since(s2)

	s3 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Sprintf("S")
	}
	t3 := time.Since(s3)

	s4 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Println("P")
	}
	t4 := time.Since(s4)

	s5 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Fprintln(w, "F")
	}
	t5 := time.Since(s5)

	s6 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Sprintln("S")
	}
	t6 := time.Since(s6)

	s7 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Print("P")
	}
	t7 := time.Since(s7)

	s8 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Fprint(w, "F")
	}
	t8 := time.Since(s8)

	s9 := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Sprint("S")
	}
	t9 := time.Since(s9)

	fmt.Printf("fmt.Printf repeat 100,000 times: %v\n", t1)
	fmt.Printf("fmt.Fprintf repeat 100,000 times: %v\n", t2)
	fmt.Printf("fmt.Sprintf repeat 100,000 times: %v\n", t3)
	fmt.Printf("fmt.Println repeat 100,000 times: %v\n", t4)
	fmt.Printf("fmt.Fprintln repeat 100,000 times: %v\n", t5)
	fmt.Printf("fmt.Sprintln repeat 100,000 times: %v\n", t6)
	fmt.Printf("fmt.Print repeat 100,000 times: %v\n", t7)
	fmt.Printf("fmt.Fprint repeat 100,000 times: %v\n", t8)
	fmt.Printf("fmt.Sprint repeat 100,000 times: %v\n", t9)

	return
}
