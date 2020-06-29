package main

import (
	"fmt"
)

type rocket int

func main() {
	var x rocket = 33
	fmt.Println(x)
	fmt.Printf("%T",x)
	var y string
	y=string(x)
	fmt.Println(y)
	fmt.Printf("%T",y)
}
