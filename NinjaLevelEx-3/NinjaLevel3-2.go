package main

import (
	"fmt"
)

func main() {
	x := 67
	fmt.Printf("%U%q\n", x, x)
	fmt.Printf("%U%q\n", x, x)
	fmt.Printf("%U%q\n\n", x, x)
	y := 68
	fmt.Printf("%U%q\n", y, y)
	fmt.Printf("%U%q\n", y, y)
	fmt.Printf("%U%q\n\n", y, y)
	z := 69
	fmt.Printf("%U%q\n", z, z)
	fmt.Printf("%U%q\n", z, z)
	fmt.Printf("%U%q\n\n", z, z)
}
