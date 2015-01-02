package main

import (
	"fmt"
)

func main() {
	var ptr *int

	var i int = 123

	ptr = &i

	fmt.Println("Address of i:", &i)
	fmt.Println("Value of ptr (address of i)", ptr)

	fmt.Println("Value of i:", i)
	fmt.Println("Value of i via pointer:", *ptr)

	*ptr = 999

	fmt.Println("Value of i via pointer", i)
}
