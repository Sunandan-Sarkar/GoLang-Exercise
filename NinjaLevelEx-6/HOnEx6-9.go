package main

import (
	"fmt"
)

func main() {
	t := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(t)
	t2 := som(t...)
	fmt.Println(t2)
	t3 := evens(som, t...)
	fmt.Println(t3)
	t4 := ods(som, t...)
	fmt.Println(t4)
}
func som(s ...int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}
func evens(f func(s ...int) int, vi ...int) int {
	var e1 []int
	for _, v := range vi {
		if v%2 == 0 {
			e1 = append(e1, v)
		}
	}
	return f(e1...)
}
func ods(f func(s ...int) int, vi ...int) int {
	var e2 []int
	for _, v := range vi {
		if v%2 != 0 {
			e2 = append(e2, v)
		}
	}
	return f(e2...)
}
