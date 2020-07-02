package main

import "fmt"

func main() {
	x1 := []int{111, 333, 444, 555}
	fmt.Println(x1)
	x3 := soo(x1...)
	fmt.Println(x3)
}
func soo(x2 ...int) int {
	fmt.Println(x2)
	tot := 0
	for _, v := range x2 {
		tot += v
	}
	return tot
}
