package main

import (
	"fmt"
)

type myData struct {
	i1, i2 int
	f      float32
	s      string
}

func main() {
	var x myData

	x.i1 = 1
	x.i2 = 2
	x.f = 3.14
	x.s = "go"

	fmt.Println(x)
}
