package main

import (
	"fmt"
)

type farmer struct {
	name     string
	age      int
	interest []string
	skill    string
	access   bool
}
type SpecialFarmer struct {
	farmer
	poor    bool
	disable bool
	sAccess string
	income  int
}

func main() {
	sf1 := SpecialFarmer{
		farmer: farmer{
			name:     "Ashraful Alam",
			age:      36,
			interest: []string{"Instagram", "TV", "Computer"},
			skill:    "Working",
			access:   false},
		poor:    true,
		disable: false,
		sAccess: "Have special access",
		income:  1000,
	}

	f1 := farmer{
		name:     "Biman Kuman",
		age:      40,
		interest: []string{"Facebook", "TV", "Radio"},
		skill:    "Farming",
		access:   true}

	fmt.Println(f1)
	fmt.Println(sf1)
	fmt.Println(f1)
	fmt.Println(sf1.sAccess, sf1.income)
	fmt.Println(sf1.disable, sf1.poor)

}
