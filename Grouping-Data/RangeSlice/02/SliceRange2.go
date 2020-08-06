package main

import "fmt"

func main() {
	y := []string{"sun", "moon", "star", "rocket", "foot", "face"}
	fmt.Println(y)
	fmt.Println(cap(y))
	fmt.Println(len(y))

	for i, v := range y {
		fmt.Println(i, v)
	}
}
