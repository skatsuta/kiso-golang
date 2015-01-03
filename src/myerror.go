package main

import (
	"fmt"
)

type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func main() {
	strs := []string{"1", "ab", "z"}

	for _, v := range strs {
		fmt.Println(hex2int(v))
	}
}

func hex2int(hex string) (val int, err error) {
	for _, r := range hex {
		val *= 16

		switch {
		case '0' <= r && r <= '9':
			val += int(r - '0')
		case 'a' <= r && r <= 'f':
			val += int(r-'a') + 10
		default:
			return 0, MyError{"invalid string: " + string(r)}
		}

	}

	return
}
