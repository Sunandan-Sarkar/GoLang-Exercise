package main

import (
	"fmt"
)

func main() {
	x := 424 //change the value of x here, same value will be printed
	if x == 40 {
		fmt.Println("our x value was 40")
	} else if x == 41 {
		fmt.Println("our x value was 41")
	} else if x == 42 {
		fmt.Println("our x value was 42")
	} else if x == 43 {
		fmt.Println("our x value was 43")
	} else {
		fmt.Println("our x value was not 40,41,42 or 43")
	}
}
