package main

import (
	"fmt"
)

func main() {
	f := func(s string) {
		fmt.Println("I Assign String value to this func", s)
	}
	f("Strange")
}
