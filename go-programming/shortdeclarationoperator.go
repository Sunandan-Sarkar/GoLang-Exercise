package main

import (
	"fmt"
)

func main() {
	x := 100 // short declaration plus assigned value
	fmt.Println(x)
	y := 67 + 9 // short declaration plus assigned value
	fmt.Println(y)
	x = 200 // only assigned value
	fmt.Println(x)
	y = 99 + 78 // only assigned value
	fmt.Println(y)
	z := "Bond, James"
	fmt.Println(z)
}
