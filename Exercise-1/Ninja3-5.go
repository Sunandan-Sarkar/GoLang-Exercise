package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("When %v is divided by 4,the modulus is %#v\n", x, x%4)
	}
}
