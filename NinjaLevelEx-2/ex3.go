package main

import (
	"fmt"
)

const (
	a int    = 42         //TYPED VARIABLE
	b        = true       //UNTYPED VARIABLE
	c string = "Novi"     //TYPED VARIABLE
	d        = 42.4647475 //UNTYPED VARIABLE
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\t", b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%T\t", d)
}
