package main

import (
	"fmt"
)

func main() {
	x := []int{3, 6, 9, 12, 15, 18}
	fmt.Println(x[5])
	fmt.Println(x)
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[:5])
	fmt.Println(x[3:4])
	fmt.Println("Here is for loop")
	for i, v := range x {
		fmt.Println(i, v)
		//Alternative method
	}
	fmt.Println("The below one is Alternative method")
	for i := 0; i <= 5; i++ {
		fmt.Println(i, x[i])
	}
}
