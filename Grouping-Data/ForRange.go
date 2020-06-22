package main

import (
	"fmt"
)

func main() {
	x := []int{3, 6, 9, 12, 15, 18}
	y := []string{"sun", "moon", "star", "rocket", "foot", "face"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(x[5])

	fmt.Println(cap(y))
	fmt.Println(y)
	fmt.Println(y[0])
	fmt.Println(y[1])
	fmt.Println(y[2])
	fmt.Println(y[3])
	fmt.Println(y[4])
	fmt.Println(y[5])
}
