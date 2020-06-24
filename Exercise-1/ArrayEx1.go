package main

import (
	"fmt"
)

func main() {

	var x [4]string//var x[5]int
	fmt.Println(x)

	x[0] = "May"
	x[1] = "June"
	x[2] = "July"
	x[3] = "August"

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
