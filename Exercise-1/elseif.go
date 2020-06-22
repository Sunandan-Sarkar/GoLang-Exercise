package main

import (
	"fmt"
)

func main() {
	x := 423
	if x == 41 {
		fmt.Println("This is 41")
	} else if x == 42 {
		fmt.Println("This is 42")
	} else if x == 43 {
		fmt.Println("This is 43")
	} else {
		fmt.Println("The value was not 40,41,42 or 43")
	}
}
