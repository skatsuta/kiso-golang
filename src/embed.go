package main

import (
	"fmt"
)

type embedded struct {
	i int
}

func (x embedded) do() {
	fmt.Println("do()")
}

type test struct {
	embedded
}

func main() {
	var x test

	x.do()

	fmt.Println(x.i)
}
