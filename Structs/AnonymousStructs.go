package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Sunandan",
		last:  "Sarkar",
		age:   33,
	}
	fmt.Println(p1)
	fmt.Println("Done.")
}
