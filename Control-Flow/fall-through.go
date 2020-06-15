package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print")
	case (2 == 2):
		fmt.Println("print")
		fallthrough
	case (4 == 4):
		fmt.Println("print?")
		fallthrough
	case (8 == 4):
		fmt.Println("This should not print")
		fallthrough
	case (42 == 42):
		fmt.Println("print true")
		fallthrough
	case (12 == 15):
		fmt.Println("Does it print?")
		fallthrough
	default:
		fmt.Println("This is default value")
	}
}
