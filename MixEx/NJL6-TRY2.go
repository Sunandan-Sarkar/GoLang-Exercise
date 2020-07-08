package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)

	ii1 := []int{11, 22, 33, 44, 55}
	n2 := bar(ii1)
	fmt.Println(n2)
}
func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
