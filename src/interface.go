package main

import (
	"fmt"
)

type Calculator interface {
	calc(x int, y int) int
}

type Add struct{}

func (Add) calc(x int, y int) int {
	return x + y
}

type Sub struct{}

func (Sub) calc(x int, y int) int {
	return x - y
}

func main() {
	var (
		add Add
		sub Sub
		cal Calculator
	)

	cal = add
	fmt.Println("Add:", cal.calc(1, 2))

	cal = sub
	fmt.Println("Sub:", cal.calc(1, 2))
}
