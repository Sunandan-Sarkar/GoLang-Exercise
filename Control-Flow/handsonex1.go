package main

import (
	"fmt"
)

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%v\t%b\t\t%#x\t%#U\n", x, x, x, x)
	}
}
