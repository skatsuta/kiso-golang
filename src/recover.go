package main

import (
	"fmt"
)

func main() {
	f(false)
	f(true)
}

func f(p bool) {
	defer func() {
		fmt.Println("defer started")

		if err := recover(); err != nil {
			fmt.Println("panic interrupted:", err)
		}

		fmt.Println("defer ended")
	}()

	if p {
		panic("panic occurred")
	}
}
