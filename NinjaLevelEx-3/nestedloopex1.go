package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ { //for init; condition; post { }
		fmt.Println("The outer loop is:", i)
		for j := 0; j < 5; j++ {
			fmt.Println("\tThe inner loop is:", j)
			for k := 0; k < 2; k++ {
				fmt.Println("\t\t The inner2 loop is:", k)
				for l := 0; l < 1; l++ {
					fmt.Println("\t\t\t The inner3 loop is:", l)
				}
			}
		}
	}
}
