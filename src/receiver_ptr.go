package main

import (
	"fmt"
)

type myType int

func (t myType) setByValue(nval myType) {
	t = nval
}

func (t *myType) setByPtr(nval myType) {
	*t = nval
}

func main() {
	var x myType = 0

	x.setByValue(1)
	fmt.Println(x)

	x.setByPtr(2)
	fmt.Println(x)
}
