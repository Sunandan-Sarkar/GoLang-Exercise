package main

import (
	"fmt"
)

func main() {
	x := "Dan"
	switch x {
	case "Novi":
		fmt.Println("This is Novi")
	case "Sun", "Dan":
		fmt.Println("This is Sun or Dan")
	case "Lim":
		fmt.Println("This is Lim")
	case "Dann":
		fmt.Println("This is Dann")
	default:
		fmt.Println("This is Default")
	}
}
