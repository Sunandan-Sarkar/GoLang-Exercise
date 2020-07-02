package main

import (
	"fmt"
)

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19}
	s := sum(xi...)
	fmt.Println("The total amount is", s)
	s2 := evn(sum, xi...)
	fmt.Println("The total amount of even numbers is", s2)
	s3 := odo(sum, xi...)
	fmt.Println("The total amount of odd numbers is", s3)
}

func sum(si ...int) int {
	total := 0
	for _, v := range si {
		total += v
	}
	return total
}

func evn(f func(si ...int) int, vi ...int) int {
	var p1 []int
	for _, v := range vi {
		if v%2 == 0 {
			p1 = append(p1, v)
		}

	}
	return f(p1...)
}

func odo(f func(si ...int) int, vi ...int) int {
	var p2 []int
	for _, v := range vi {
		if v%2 != 0 {
			p2 = append(p2, v)
		}
	}
	return f(p2...)
}
