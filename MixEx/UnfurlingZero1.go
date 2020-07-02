package main

import (
	"fmt"
)

func main() {
	p := sum()
	fmt.Println(p)

}
func sum(x ...int) int {
	fmt.Println(x)
	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, v, sum)
	}
	return sum

}
