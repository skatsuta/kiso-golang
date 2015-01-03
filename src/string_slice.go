package main

import (
	"fmt"
)

func main() {
	var x string = "abcde"[1:4]

	var s string = "あいうえお"
	var y string = s[3:9]
	var z string = s[1:4]

	fmt.Println(x, y, z)
}
