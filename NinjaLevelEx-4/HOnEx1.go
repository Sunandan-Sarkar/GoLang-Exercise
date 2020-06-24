package main

import (
	"fmt"
)

func main() {
	//create an ARRAY which holds 5 VALUES of TYPE int
	var x [5]int
	fmt.Println(x)
	fmt.Println("Hello, Sunandan")
	//assign VALUES to each index position.
	x[0] = 42
	x[1] = 43
	x[2] = 44
	x[3] = 45
	x[4] = 46
	fmt.Println(x)
	fmt.Println("Hello, Sunandan")
	//Range over the array and print the values out.

	for i, v := range x {
		fmt.Println(i, v)

	}
	//Using format printing, print out the TYPE of the array
	fmt.Printf("%T\n", x)
	fmt.Println("Hello, Sunandan")

	//x := [5]int{11, 22, 33, 44, 55}
	//	fmt.Println(x)
	//Alternative method
	//	for i, v := range x {
	//		fmt.Println(i, v)
	//	}
	//	fmt.Printf("%T\n", x)

}
