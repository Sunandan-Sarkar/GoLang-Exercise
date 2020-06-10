package main

import (
	"fmt"
)

const (
	a = 42
	b = 45.56789
	c = "Sunandan Sarkar"
)

//Or separately
//const a = 42
//const b = 45.56789
//const c = "Sunandan Sarkar"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
