package main

import (
	"fmt"
)

func main() {
	// make a channel typed int
	c := make(chan int)

	go func(s chan<- int) {
		for i := 0; i < 5; i++ {
			s <- i
		}

		close(s)
	}(c)

	for {
		val, open := <-c

		if !open {
			break
		}

		fmt.Println(val)
	}
}
