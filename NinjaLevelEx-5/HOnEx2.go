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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for _, v := range m {
		fmt.Println(v.first, v.last)
		for k, l := range v.iceflavors {
			fmt.Println(k, l)
		}
		fmt.Println("-------")
	}

}
