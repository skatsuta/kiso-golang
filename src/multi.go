package main

import (
	"fmt"
)

func calc(a int, b int) (int, int, int, float32) {
	return a + b, a - b, a * b, float32(a) / float32(b)
}

func main() {
	fmt.Println(calc(1, 2))
}
