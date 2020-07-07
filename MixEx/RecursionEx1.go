package main

import (
	"fmt"
)

func main() {
	fmt.Println()
	x := factorial(4)
	fmt.Println(x)

	x2 := loopfact(4)
	fmt.Println(x2)
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func loopfact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
