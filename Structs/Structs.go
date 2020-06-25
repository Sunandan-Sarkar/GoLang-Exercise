package main

import (
	"fmt"
)

type Bank struct {
	account string
	number  int
	valid   bool
	date    []string
	age     int
}

func main() {
	p1 := Bank{
		account: "Sunandan Sarkar",
		number:  1234567890,
		valid:   true,
		date:    []string{"June 25th,2020"},
		age:     34,
	}
	p2 := Bank{
		account: "Limbaroyati",
		number:  9876543210,
		valid:   true,
		date:    []string{"June 26th,2020"},
		age:     28,
	}
	fmt.Println(p1)
	fmt.Println(p1.account, p1.number, p1.valid, p1.date, p1.age)
	fmt.Println(p2)
	fmt.Println(p1.account)
	fmt.Println(p1.number)
	fmt.Println(p1.valid)
	fmt.Println(p1.date)
	fmt.Println(p1.age)
	fmt.Println(p2.account, p2.number, p2.valid, p2.date, p2.age)
	fmt.Println(p2.account)
	fmt.Println(p2.number)
	fmt.Println(p2.valid)
	fmt.Println(p1.date)
	fmt.Println(p2.age)

}
