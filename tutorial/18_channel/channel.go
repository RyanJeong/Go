package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := range ch { // range ch = <-ch
			fmt.Println("received:", i)
			//fmt.Println("received:", <-ch)
		}
	}()
	ch <- 1
	close(ch)
}
