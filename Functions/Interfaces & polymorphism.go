package main

import (
	"fmt"
)

type farmer struct {
	fname string
	lname string
	age   int
}
type Sfarmer struct {
	farmer
	income   []string
	landhold bool
}

func main() {
	f1 := farmer{
		fname: "Biman",
		lname: "Kumar",
		age:   42,
	}
	f2 := Sfarmer{
		farmer: farmer{
			fname: "Ashraful",
			lname: "Alam",
			age:   35,
		},
		income:   []string{"Monthly income is ", "300 euro"},
		landhold: true,
	}
	fmt.Println(f1)
	fmt.Println(f2)
	f1.foo()
	f2.foo()
	bar(f1)
	bar(f2)
}
func (fa1 farmer) foo() {
	fmt.Println(fa1.fname, fa1.lname)
}
func (fa2 Sfarmer) foo() {
	fmt.Println(fa2.fname, fa2.lname, fa2.landhold, fa2.income)
	fmt.Println("Polymorphism")
}

type human interface {
	foo()
}

func bar(h human) {
	fmt.Println(h)
}
