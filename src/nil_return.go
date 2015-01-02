package main

import (
	"fmt"
)

type MyError string

func (e MyError) Error() string {
	return string(e)
}

func do() error {
	var ret *MyError = nil

	return ret
}

func main() {
	if err := do(); err != nil {
		fmt.Printf("Error exists: %T is not %T", err, nil)
	} else {
		fmt.Println("No error")
	}
}
