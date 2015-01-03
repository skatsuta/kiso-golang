package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main: started")

	fmt.Println("call test() as a usual func")
	test()

	fmt.Println("call test() as a goroutine")
	go test()

	time.Sleep(3 * time.Second)

	fmt.Println("main: exited")
}

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		time.Sleep(1 * time.Second)
	}
}
