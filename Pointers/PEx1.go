package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // gives the address
	fmt.Println(a)

	fmt.Println("---Below the type---")

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	fmt.Println("---Below---")
	b := &a
	fmt.Println(b)
	fmt.Println(*b) //gives the value of the address

	fmt.Println("---Below the new assigned value---")

	*b = 50
	fmt.Println(a)
	fmt.Println(*b)
}
