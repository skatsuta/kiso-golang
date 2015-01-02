package main

import (
	"fmt"
)

func main() {
	var i interface{} = "test"

	s, ok := i.(string)

	fmt.Println(s, ok)

	v, ok := i.(int)

	fmt.Println(v, ok)
}
