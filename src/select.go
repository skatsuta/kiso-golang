package main

import (
	"fmt"
	"os"
)

func main() {
	c1 := make(chan int, 1)
	c2 := make(chan string, 1)

	go func() {
		for i := 0; i < 5; i++ {
			select {
			case c1 <- 0:
				fmt.Println("send to ch 1")
			case c2 <- "test":
				fmt.Println("send to ch 2")
			}
		}

		os.Exit(0)
	}()

	for {
		select {
		case v := <-c1:
			fmt.Println("receive from ch 1:", v)
		case v := <-c2:
			fmt.Println("receive from ch 2:", v)
		}
	}
}
