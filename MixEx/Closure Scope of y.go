package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println(x)
	x++
	{
		y := 42
		fmt.Println(y)
	}
	//error happens because y inside the block
	//fmt.Println(y)

	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
}
func foo() {
	fmt.Println("Tere!")
}
