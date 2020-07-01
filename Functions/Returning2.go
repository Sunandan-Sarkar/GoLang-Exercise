package main

import (
	"fmt"
)

func main() {
	s1 := foo()
	fmt.Println(s1)
	b1 := boo()
	fmt.Println(b1)
	c1 := coo()
	fmt.Println(c1)
}
func foo() string {
	s := "Hello=====Tallinn"
	return s
}
func boo() int {
	b := 333
	return b
}
func coo() string {
	c := ("Sunandan Sarkar")
	return c
}
