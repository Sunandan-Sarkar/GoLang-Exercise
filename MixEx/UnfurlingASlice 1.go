package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 3, 5, 7, 9}
	s := sum(xi...)
	fmt.Println("Total", s)
}

func sum(x ...int) int {
	fmt.Println(x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("In position", i, "now adding", v, "Total became", sum)
	}
	return sum
}
