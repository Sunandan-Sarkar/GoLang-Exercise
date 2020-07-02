package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)

}
func sum(s ...int) {
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	sum := 0
	for i, j := range s {
		sum += j
		fmt.Println("item in position", i, "now adding", j, "total become", sum)
	}
	fmt.Println("Total is:", sum)
}
