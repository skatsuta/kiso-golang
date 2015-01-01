package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j float64 = 4, 16

	var op string
	x, y := math.Sqrt(i), math.Sqrt(j)

	if x < y {
		op = "<"
	} else if x > y {
		op = ">"
	} else {
		op = "="
	}

	fmt.Println(x, op, y)
}
