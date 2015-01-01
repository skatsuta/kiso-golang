package main

import (
	"fmt"
)

func main() {
	s := "abcあいう"

	for i, u := range s {
		fmt.Println(i, u)
	}
}
