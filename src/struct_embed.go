package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
	Person
}

func main() {
	e := &Employee{
		id: 1,
		Person: Person{
			name: "Jack",
			age:  28,
		},
	}

	fmt.Printf("%v\n%+v\n%#v", e, e, e)
}
