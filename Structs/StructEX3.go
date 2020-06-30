package main

import (
	"fmt"
)

func main() {
	f1 := struct {
		name     string
		age      int
		interest []string
		skill    string
		access   bool
	}{
		name:     "Biman Kuman",
		age:      40,
		interest: []string{"Facebook", "TV", "Radio"},
		skill:    "Farming",
		access:   true}

	fmt.Println(f1)
	fmt.Println(f1.skill)
	fmt.Println(f1.access)
	fmt.Println(f1.age)

}
