package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	goroutines = 10
	maxProcs   = 3
)

func main() {
	semaphore := make(chan int, maxProcs)

	notify := make(chan int, goroutines)

	for i := 0; i < goroutines; i++ {
		go func(no int, semaphore chan int, notify chan<- int) {
			// 自分の番が来るのを待つ
			semaphore <- 0

			time.Sleep(time.Duration(rand.Int63n(3)) * time.Second)

			fmt.Println("process finished", no)

			// 待機中の goroutine に処理を明け渡す
			<-semaphore

			notify <- no
		}(i, semaphore, notify)
	}

	for i := 0; i < goroutines; i++ {
		<-notify
	}

	fmt.Println("all finished")
}
