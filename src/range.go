package main

import (
	"fmt"
)

func main() {
	arr := []string{"Alice", "Bob", "Cott"}

	for i, e := range arr {
		fmt.Printf("%v: %v\n", i, e)
	}
}
