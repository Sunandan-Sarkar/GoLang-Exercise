package main

import (
	"fmt"
)

func main() {
	for p := 100; p <= 900; p++ {
		fmt.Printf("When %v is divided by 16, The modulus is %#v\n", p, p%16)
	}
}
