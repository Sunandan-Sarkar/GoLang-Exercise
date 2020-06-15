package main

import (
	"fmt"
)

func main() {
	// or can define a variable n:="Bond", switch n
	switch "Bond" {
	case "honeybunny":
		fmt.Println("This is honeybunny")
	case "Bond":
		fmt.Println("This is bond")
	case "Novi":
		fmt.Println("This is Novi")
	default:
		fmt.Println("This is default value")
	}
}
