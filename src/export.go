package main

import (
	"fmt"
)

const (
	Export = true
	export = false
)

func main() {
	const (
		Z = 123
	)

	fmt.Println(export, Z)
}
