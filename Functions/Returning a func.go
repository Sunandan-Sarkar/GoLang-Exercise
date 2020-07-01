package main

import "fmt"

func main() {
	foo(2020) // usual function

	b1 := bar("Print")
	fmt.Println(b1)
}

//usual function
func foo(x int) {
	fmt.Println(x)
}

//Returning
func bar(b string) string {
	return b
}
