package main

import (
	"fmt"
)

func main() {
	type myInt int
	var i myInt = 123
	i += 1
	fmt.Println(i)

	type myStruct struct {
		a int
		b int
	}
}
