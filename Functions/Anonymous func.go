package main

import (
	"fmt"
)

func main() {
	foo() //usual function

	//Anonymous Function
	func() {
		fmt.Println()
	}() //() no passing arguments
	func(x int) {
		fmt.Println("This is Anonymous Function with int value", x)
	}(33)
	func(y string) {
		fmt.Println("This is Anonymous Function with string value:", y)
	}("I am string")
}

//usual function
func foo() {
	fmt.Println("This is usual function.")
}
