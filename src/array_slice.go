package main

import (
	"fmt"
)

func double(vals []int) {
	for i, l := 0, len(vals); i < l; i++ {
		vals[i] *= 2
	}
}

func main() {
	vals := [...]int{0, 1, 2, 3}

	double(vals[:])

	fmt.Println(vals)
}
