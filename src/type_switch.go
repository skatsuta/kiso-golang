package main

import (
	"fmt"
)

func main() {
	for _, v := range typestr(nil, 123, "abc", 3.14) {
		fmt.Println(v)
	}
}

func typestr(items ...interface{}) (types []string) {
	types = make([]string, len(items))
	fmt.Printf("size of %v is %v\n", types, len(types))

	for i, v := range items {
		types[i] = getType(v)
	}

	return
}

func getType(x interface{}) string {
	switch x.(type) {
	case nil:
		return "nil"
	case int, int32, int64:
		return "integer"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
