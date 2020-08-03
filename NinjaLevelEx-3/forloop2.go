package main

import (
	"fmt"
)

func main() { //for { }
	p := 0
	for {
		if p > 5 { // i feel if statement is a condition
			break
		}
		fmt.Println(p)
		p++
	}
}
