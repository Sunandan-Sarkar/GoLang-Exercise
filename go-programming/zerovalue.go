package main

import (
	"fmt"
)

var z string = "I am in Tallinn"
var a int = 99

func main() {
	fmt.Println(z, "why? Because I study here")
	fmt.Printf("%T\n", z)

	fmt.Println("I am", a, "years old man")
	fmt.Printf("%T\n", a)

}
