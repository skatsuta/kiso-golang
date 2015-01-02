package main

import (
	"fmt"
)

type Person struct {
	name string
	age  uint
}

func main() {
	var john Person

	john.name = "John"
	john.age = 23

	tom := Person{
		age:  31,
		name: "Tom",
	}

	jane := Person{"Jane", 42}

	mike := &Person{
		name: "Mike",
		age:  36,
	}

	fmt.Println(john, tom, jane, mike)
}
