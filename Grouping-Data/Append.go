package main

import (
	"fmt"
)

func main() {
	x := []string{"Sun", "Moon", "Star", "Jupitar", "Earth", "Planet"}
	fmt.Println(x)
	x = append(x, "Tree", "Forest", "Land")
	fmt.Println(x)

	y := []string{"Honey", "Penny", "Money", "Funny"}
	x = append(x, y...)
	fmt.Println(x)
}
