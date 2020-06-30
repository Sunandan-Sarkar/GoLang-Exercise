package main

import (
	"fmt"
)

type animal struct {
	name   string
	origin string
	price  int
}
type Sanimal struct {
	animal
	special bool
	remarks []string
}
type Smanimal struct {
	Sanimal
	sales     int
	available bool
}
type allanimal interface {
	loo()
}

func main() {
	a1 := animal{
		name:   "Bengal Tiger",
		origin: "Bangladesh",
		price:  100000,
	}
	a2 := Sanimal{
		animal: animal{
			name:   "Dog",
			origin: "Japan",
			price:  25000,
		},
		special: true,
		remarks: []string{"Strong", "Friendly", "Danger to stranger"},
	}
	a3 := Sanimal{
		animal: animal{
			name:   "Cat",
			origin: "England",
			price:  5000,
		},
		special: false,
		remarks: []string{"Fair", "Friendly", "Silent killer"},
	}
	a4 := Smanimal{
		Sanimal: Sanimal{
			animal: animal{
				name:   "Deer",
				origin: "Estonia",
				price:  3500,
			},
			special: true,
			remarks: []string{"Awesome", "Beautiful", "illus"}},
		sales:     999,
		available: false,
	}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	a1.loo()
	a2.loo()
	a3.loo()
	a2.loo()
	a3.loo()
	a4.loo()
	loo(a1)
	loo(a2)
	loo(a3)
	loo(a4)
}
func (aa1 animal) loo() {
	fmt.Println("My", aa1)
	fmt.Println("My", aa1.name, aa1.price, aa1.origin)
}
func (aa2 Sanimal) loo() {
	fmt.Println("My", aa2)
	fmt.Println("---------------------")
}
func (aa3 Smanimal) loo() {
	fmt.Println("My dear.....", aa3)
}
func loo(ap allanimal) {
	switch ap.(type) {
	case animal:
		fmt.Println(ap, "is My favourite")
	case Sanimal:
		fmt.Println(ap, "is My favourite")
	case Smanimal:
		fmt.Println(ap, "is My favourite")
	}
	fmt.Println("My favourite", ap)
}
