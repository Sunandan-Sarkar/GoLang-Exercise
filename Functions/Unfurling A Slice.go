package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	s := sum(xi...)
	fmt.Println("The total amount is:", s)
}
func sum(s ...int) int {
	fmt.Println(s)
	sum := 0
	for i, v := range s {
		sum += v
		fmt.Println("Index position number:", i, "Now adding ", v, "Total became: ", sum)
	}
	fmt.Println("The total amount is:",sum)
	return sum
}
