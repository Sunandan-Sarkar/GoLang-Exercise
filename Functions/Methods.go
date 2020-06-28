package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}
type Sperson struct {
	person
	access bool
}

func main() {
	sp1 := Sperson{
		person: person{
			first: "Ashraful",
			last:  "Alam",
		},
		access: true,
	}
	sp2 := Sperson{
		person: person{
			first: "Biman",
			last:  "Kumar",
		},
		access: false,
	}
	fmt.Println(sp1)
	fmt.Println(sp2)

	sp1.speak()
	sp2.speak()
}
func (s Sperson) speak() {
	fmt.Println("I am", s.first, s.last)
}
