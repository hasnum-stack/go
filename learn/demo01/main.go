package main

import (
	"demo01/gotest"
	"fmt"
)

type fn1 func(int) int

func test(a int, b string) fn1 {
	return func(i int) int {
		return i
	}
}

func main() {
	a := 1
	fmt.Printf("a: %v\n", a)
	r := test(1, "123")
	fmt.Printf("r: %v\n", r)
	gotest.TestTest()
	test5()
}
