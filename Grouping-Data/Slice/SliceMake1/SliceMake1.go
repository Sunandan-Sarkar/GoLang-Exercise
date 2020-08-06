package main

import (
	"fmt"
)

func main() {
	x := []int{3, 6, 9, 12, 15, 18}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[2] = 911
	x[9] = 120
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 1000)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 98765)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 16740)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 9747447, 98669, 325463, 2566, 578, 68070, 7909, 5363632, 787, 679, 675745)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
