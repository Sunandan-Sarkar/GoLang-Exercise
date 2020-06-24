package main

import (
	"fmt"
)

func main() {
	//Using a COMPOSITE LITERAL, create a SLICE of TYPE int, assign 10 VALUES

	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(x)
	//Range over the slice and print the values out.

	for i, v := range x {
		fmt.Println(i, v)

	}
	//Using format printing,print out the TYPE of the slice
	fmt.Printf("%T\n", x)
}
