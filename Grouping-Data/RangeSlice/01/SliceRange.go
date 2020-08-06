package main

import (
	"fmt"
)

func main() {
	x := []int{3, 6, 9, 12, 15, 18}
	fmt.Println(x)
	fmt.Println(len(x))
	for i, v := range x {
		fmt.Println(i, v)
	}

}
