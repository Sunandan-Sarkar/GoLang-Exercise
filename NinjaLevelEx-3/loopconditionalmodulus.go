package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("Prints")
	case true:
		fmt.Println("Don't print")
		fallthrough
	case (3 != 4):
		fmt.Println("Does it print")
		fallthrough
	case (4 == 4):
		fmt.Println("This should print")
		fallthrough
	default:
		fmt.Println("This is default value")
	}
}
