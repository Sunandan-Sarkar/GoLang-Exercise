package main

import (
	"fmt"
)

func main() {
	n := "Bond"
	switch n {
	case "honeybunny", "bond", "Novi":
		fmt.Println("This is honeybunny or Bond or Novi")
	case "B":
		fmt.Println("This is bond")
	case "N":
		fmt.Println("This is Novi")
	default:
		fmt.Println("\t\tThis is default value")
	}
}
