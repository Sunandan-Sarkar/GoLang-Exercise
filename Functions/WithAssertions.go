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

//Methods
func (fa1 farmer) foo() {
	fmt.Println(fa1.fname, fa1.lname, "-is the farmer")
}

//Methods
func (fa2 Sfarmer) foo() {
	fmt.Println(fa2.fname, fa2.lname, "-is the farmer")
}

//Methods
func (fa3 Sfarmer) boo() {
	fmt.Println(fa3.fname, fa3.lname, "-is the farmer from boo method")

	fmt.Println("---Below is the Polymorphism---")
}

type human interface {
	foo()
}

//functions
func bar(h human) {
	switch h.(type) {
	case farmer:
		fmt.Println(h.(farmer).fname, " is also called human in switchcase")
	case Sfarmer:
		fmt.Println(h.(Sfarmer).fname, " is also called human in switchcase")
	}
	fmt.Println(h, " is also called human")
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
	f3 := Sfarmer{
		farmer: farmer{
			fname: "Fatima",
			lname: "Begam",
			age:   30,
		},
		income:   []string{"Monthly income is ", "200 euro"},
		landhold: true,
	}
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	f1.foo()
	f2.foo()
	f3.boo()
	bar(f1)
	bar(f2)
	bar(f3)
}
