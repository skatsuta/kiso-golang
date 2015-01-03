package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)

	c <- 0

	fmt.Printf("%#v, len: %v, cap: %v\n", c, len(c), cap(c))
}
