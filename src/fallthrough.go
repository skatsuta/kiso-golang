package main

import (
	"fmt"
	"time"
)

func main() {
	for day := time.Sunday; day <= time.Saturday; day++ {
		switch day {
		case time.Sunday:
			fallthrough
		case time.Saturday:
			fmt.Println(day, "holiday")
		default:
			fmt.Println(day, "maybe weekday")
		}
	}
}
