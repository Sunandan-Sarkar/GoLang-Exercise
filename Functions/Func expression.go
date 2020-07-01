package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("This is func expression without passing arguments")
	}
	f()
	g := func(x int) {
		fmt.Println("This is func expression with passing arguments int", x)
	}
	g(334455)
	h := func(x string) {
		fmt.Println("This is func expression with passing arguments string:", x)
	}
	h("This is String value")

}
