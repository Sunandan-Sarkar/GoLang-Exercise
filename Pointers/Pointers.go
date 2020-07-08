package main

import (
	"fmt"
)

func main() {
	x := 99
	fmt.Println(x)
	fmt.Println(&x) // & gives the address

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", &x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y) // * gives the value stored at an address when you have the address
	fmt.Println(&x)

	*y = 100
	fmt.Println(x)
	fmt.Println(&x)
}
