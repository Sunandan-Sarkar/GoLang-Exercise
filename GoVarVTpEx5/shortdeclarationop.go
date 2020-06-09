package main

import (
	"fmt"
)

var a int
var i string = "Novi"
var b string = "Be Responsible"
var c bool
var d bool = true

func main() {
	e := 98
	f := "Shaken not stirred"
	g := `Limbaroyati says "Hello, Sunandan!"`
	h := `Sunandan replied
"Hello Limbaroyati! What's up?'"`
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(i, "Says", f)
	fmt.Println(i, "Says", b)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
