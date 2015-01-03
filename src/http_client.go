package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "http://golang.jp/hello.html"

	res, err := http.Get(url)

	check(err)

	body, err := ioutil.ReadAll(res.Body)

	check(err)

	fmt.Println("status:", res.Status)
	fmt.Println("content-type:", (res.Header["Content-Type"])[0])
	fmt.Println(string(body))
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
