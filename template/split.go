package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := strings.Split(time.Now().Format("2006-01-02"), "-")
	fmt.Printf("%s\n%s\n%s", s[0], s[1], s[2])
}
