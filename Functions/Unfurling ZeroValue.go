package main

import (
	"fmt"
)

func main() {
	s := sum()
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
