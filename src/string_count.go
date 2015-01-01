package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var ja string = "Go言語"

	fmt.Println(ja, "count:", utf8.RuneCountInString(ja))
}
