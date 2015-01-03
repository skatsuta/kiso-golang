package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 10, 20)
	s2 := make([]float32, 5)

	fmt.Printf("%#v, len: %v, cap: %v\n", s1, len(s1), cap(s1))
	fmt.Printf("%#v, len: %v, cap: %v\n", s2, len(s2), cap(s2))

}
