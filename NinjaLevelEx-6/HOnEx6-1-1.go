package main

import (
	"fmt"
)

func main() {
	n := foo()
	fmt.Println(n)
	m, p := bar()
	fmt.Println(m, p)
}
func foo() int {
	return 44
}
func bar() (int, string) {
	return 99, "Sundan"
}
