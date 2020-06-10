package main

import (
	"fmt"
)

var x int
var y float64

func main() {
	//	x:=42
	//	y:=42.768765 or below assigned way
	x = 42
	y = 42.6566555
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
