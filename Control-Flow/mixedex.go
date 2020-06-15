package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 9; i++ {
		fmt.Printf("The outer loop:%d\n", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("\t\tThe inner Loop:%d\n", j)
		}
	}
}
