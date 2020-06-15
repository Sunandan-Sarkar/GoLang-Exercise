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
	case (4 == 4):
		fmt.Println("Does it print?")
	}
}
