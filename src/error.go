package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
	}

	fmt.Println("OK")
}
