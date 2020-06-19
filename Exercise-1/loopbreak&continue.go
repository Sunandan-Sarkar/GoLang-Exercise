package main

import (
	"fmt"
)

func main() {
	x := 0
	for {
		if x > 10 {
			break
		}
		fmt.Println(x)
		x++
	}
}
