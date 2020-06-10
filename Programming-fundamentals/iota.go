package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
	d = iota
	e = iota
	f = iota
)

// Reset and start new
//const (
//	d = iota
//	e = iota
//	f = iota
//)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T %T %T %T %T %T", a, b, c, d, e, f)
}
