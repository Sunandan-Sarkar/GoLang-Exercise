package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Queen", "Zoo", "Alice", "Dr. Strange"}

	fmt.Println("Befor", xi)
	sort.Ints(xi)
	fmt.Println("After", xi)

	fmt.Println("Befor", xs)
	sort.Strings(xs)
	fmt.Println("After", xs)
}
