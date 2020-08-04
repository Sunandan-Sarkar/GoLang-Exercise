package main

import "fmt"

func main() {
	fmt.Println(" 3 + 7 =", myTot(3, 7))
	fmt.Println(" 13 + 22 =", myTot(13, 22))
	fmt.Println(" 23 + 17 =", myTot(23, 17))

}
func myTot(xi ...int) int {
	tot := 0
	for _, v := range xi {
		tot += v
	}
	return tot
}
