package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	filename = "test.txt"
)

func main() {
	write()
	read()
}

func write() {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString("metaprogramming ruby\n")

	fmt.Fprintf(file, "angularjs reference\n")
}

func read() {
	fmt.Println("=== read by using bufio.Scanner()")

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\n=== read by using ioutil.ReadFile()")

	dat, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(string(dat))
}
