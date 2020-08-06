package main

import (
	"fmt"
)

func main() {
	x := []string{"Sun", "Moon", "Star", "Jupitar", "Earth", "Planet"}

	fmt.Println(x)

	y := []string{"Honey", "Penny", "Money", "Funny", "Cunny", "Sunny"}

	fmt.Println(y)
	fmt.Println(` New line
	 			new line 
		new line
		
		`)

	mp := [][]string{x, y}
	fmt.Println(mp)
}
