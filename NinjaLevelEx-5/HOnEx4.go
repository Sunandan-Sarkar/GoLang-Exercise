package main

import (
	"fmt"
)

func main() {
	s1 := struct {
		name     string
		sex      string
		location string
		age      int
		id       int
	}{
		name:     "Sunandan Sarkar",
		sex:      "Male",
		location: "Tallinn",
		age:      34,
		id:       228287,
	}
	fmt.Println(s1)
	fmt.Println(s1.sex,s1.id)
}
