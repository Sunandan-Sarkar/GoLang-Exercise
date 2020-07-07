package main

import (
	"fmt"
)

func main() {
	fmt.Println("Create a func with the identifier foo that returns an int")

	s1 := foo()
	fmt.Println(s1)

	s2, s3 := bar()
	fmt.Println(s2, s3)
}
func foo() int {
	n := 42
	return n
}

//create a func with the identifier foo that returns an int

//create a func with the identifier bar that returns an int and a string

func bar() (int, string) {
	a := 77
	b := "Sunandan"
	return a, b
}
