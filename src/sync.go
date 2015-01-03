package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	goroutines = 3
)

func main() {
	c := make(chan int)

	for i := 0; i < goroutines; i++ {
		go func(s chan<- int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			fmt.Println("process completed")

			s <- 0
		}(c)
	}

	for i := 0; i < goroutines; i++ {
		<-c
	}

	fmt.Println("all completed")
}
