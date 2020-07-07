package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sl)
	sl2 := foo(sl...)
	fmt.Println(sl2)

	sl3 := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	sl4 := bar(sl3)
	fmt.Println(sl4)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
