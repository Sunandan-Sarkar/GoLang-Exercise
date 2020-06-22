package main

import (
	"fmt"
)

func main() {
	//x:= type{value} Composite literal
	x := []int{0, 1, 2, 3, 42, 8}
	y := []string{"sun", "moon", "star", "rocket", "foot", "face"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(y)
	fmt.Println(len(y))

}
