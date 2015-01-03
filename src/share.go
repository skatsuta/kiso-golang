package main

import (
	"fmt"
	"os"
)

const (
	goroutines = 10
)

func main() {
	counter := make(chan int, 1)

	for i := 0; i < goroutines; i++ {
		go func(counter chan int) {
			val := <-counter

			val++

			fmt.Println("counter:", val)

			if val == goroutines {
				os.Exit(0)
			}

			counter <- val
		}(counter)
	}

	counter <- 0

	for {
	}
}
