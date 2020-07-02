package main

import (
	"fmt"
)

func main() {
	w := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(w)
	w2 := one(w...)
	fmt.Println("This is the all numbers total", w2)
	w3 := even(one, w...)
	fmt.Println("This is the even numbers total", w3)
	w4 := odd(one, w...)
	fmt.Println("This is the odd numbers total", w4)
}
func one(on ...int) int {
	//fmt.Println(on)
	n := 0
	for _, v := range on {
		n += v
	}
	return n
}
func even(f func(on ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}
func odd(f func(on ...int) int, vi ...int) int {
	var y2 []int
	for _, v := range vi {
		if v%2 != 0 {
			y2 = append(y2, v)
		}
	}
	return f(y2...)
}
