package main

import (
	"fmt"
)

func main() {
	w := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(w)
	w1 := ioo(w...)
	fmt.Println(w1)

	w2 := evensum(ioo, w...)
	fmt.Println(w2)
	w3 := odd(ioo, w...)
	fmt.Println(w3)
}

func ioo(s ...int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}
func evensum(f func(s ...int) int, vi ...int) int {
	var es []int
	for _, v := range vi {
		if v%2 == 0 {
			es = append(es, v)
		}
	}
	return f(es...)
}
func odd(f func(s ...int) int, vi ...int) int {
	var es1 []int
	for _, v := range vi {
		if v%2 != 0 {
			es1 = append(es1, v)
		}
	}
	return f(es1...)
}
