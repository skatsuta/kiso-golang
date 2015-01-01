package main

import (
	"fmt"
)

const (
	// 型をもたない定数
	C1 = 12345
	// 型 (int) をもつ定数
	C2 int = 23456
)

func main() {
	// C2 が int なので、型をもたない C1 も int として計算される
	var x int = C1 * C2

	fmt.Println("12345 * 23456 =", x)

	var a int32 = 123
	var b int64 = 234

	fmt.Println("123 + 234 =", a+int32(b))
}
