package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer")

	f()
}

func f() {
	panic("panic occurred")
}
