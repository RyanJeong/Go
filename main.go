package main

import "fmt"

func main() {
	i := 1
	ch := make(chan int)
	ch <- i
	j := <-ch
	fmt.Println(j)
}
