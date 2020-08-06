package main

import (
	"fmt"
)

func main() {
	x := []string{"Sun", "Moon", "Star", "Jupitar", "Earth", "Planet"}
	fmt.Println(x)
	x = append(x, "Tree", "Forest", "Land") //func append(slice []T, elements ...T) []T
	fmt.Println(x)

	y := []string{"Honey", "Penny", "Money", "Funny"}
	x = append(x, y...) //x = append(x, 4, 5, 6)
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println("\tAnother one")
	for i := 0; i <= 12; i++ {
		fmt.Println(i, x[i])
	}
	//delete Jupitar and Earth
	x = append(x[:2], x[5:]...)
	fmt.Println(x)

}
