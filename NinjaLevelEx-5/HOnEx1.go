package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	iceflavors []string
}

func main() {
	p1 := person{
		first:      "Rocket",
		last:       "Man",
		iceflavors: []string{"Vanilla", "Hazelnut", "Spicy"},
	}
	p2 := person{
		first:      "Covid",
		last:       "Man19",
		iceflavors: []string{"Chocolate", "Hot", "Strawberry"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p2.first)
	fmt.Println(p1.last, p2.last)
	fmt.Println(p1.iceflavors, p2.iceflavors)
	fmt.Println(p1.first, p1.last)
	for i, v := range p1.iceflavors {
		fmt.Println("\t", i, v)
	}
	fmt.Println(p2.first, p2.last)
	for j, k := range p2.iceflavors {
		fmt.Println("\t", j, k)
	}

}
