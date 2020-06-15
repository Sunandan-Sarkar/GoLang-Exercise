package main

import "fmt"

func main() {
	//x := 83 / 40 // remainder
	//y := 2 % 2 // modulus
	//fmt.Println(x, y)// result is: 2,0
	x := 0
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)

	}

}
