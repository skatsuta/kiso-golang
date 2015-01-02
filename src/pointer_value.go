package main

import (
	"fmt"
)

func main() {
	var a, b int = 1, 1

	var i *int = new(int)
	var s *string = new(string)

	// a is passed by value
	// b is passed by reference
	double(a, &b)

	fmt.Println("pass by value:", a)
	fmt.Println("pass by reference:", b)

	fmt.Println(*i, *s)
}

func double(x int, y *int) {
	x *= 2
	*y *= 2
}
