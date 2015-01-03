package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	host := "www.google.com"

	addrs, err := net.LookupHost(host)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(strings.Join(addrs, "\n"))
}
