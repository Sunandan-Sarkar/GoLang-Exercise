package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
	t1 := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	fmt.Println(t1)
	t2 := too(t1...)
	fmt.Println(t2)
}
func too(sl ...int) int {
	fmt.Println(sl)
	sm := 0
	for _, v := range sl {
		sm += v
	}
	return sm
}
