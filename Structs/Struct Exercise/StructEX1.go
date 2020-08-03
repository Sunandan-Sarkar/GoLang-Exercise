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

func main() {
	f1 := farmer{
		name:     "Biman Kuman",
		age:      40,
		interest: []string{"Facebook", "TV", "Radio"},
		skill:    "Farming",
		access:   true,
	}

	f2 := farmer{
		name:     "Ashraful Alam",
		age:      36,
		interest: []string{"Instagram", "TV", "Computer"},
		skill:    "Working",
		access:   false,
	}
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f1.age, f2.age)
	fmt.Println(f1.skill, f2.skill)
	fmt.Println(f1.access, f2.access)
	fmt.Println(f1.name, f2.name)
}
