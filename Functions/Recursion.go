package main

import (
	"fmt"
)

func main() {
	p := Factloop(4)
	fmt.Println(p)
}
func Factloop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
