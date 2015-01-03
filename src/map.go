package main

import (
	"fmt"
)

func main() {
	capitals := make(map[string]string, 3)

	capitals["Japan"] = "Tokyo"
	capitals["US"] = "Washington D. C."
	capitals["China"] = "Beijing"

	for k, v := range capitals {
		fmt.Printf("%v: %v\n", k, v)
	}

	// キーの存在確認
	capital, ok := capitals["UK"]
	if ok {
		fmt.Println("registered", capital)
	} else {
		fmt.Println("unregistered")
	}
}
