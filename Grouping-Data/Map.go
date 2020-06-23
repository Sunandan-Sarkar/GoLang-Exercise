package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Forest Gump":                1994,
		"The Greenbook":              2018,
		"Departed":                   2006,
		"Inception":                  2010,
		"The Good,the Bad, the Ugly": 1966,
		"Crash":                      2004,
	}
	fmt.Println(m)
	fmt.Println(m["Inception"])
	fmt.Println(m["Cast away"])
	fmt.Println(m["Departed"])

	v, ok := m["Forest Gump"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["The Good,the Bad, the Ugly"]; ok {
		fmt.Println(v)
	}
	if p, ok := m["Departed"]; ok {
		fmt.Println(p)
	}
	fmt.Println(m["Crash out"])

	if p, ok := m["Crash"]; ok {
		fmt.Println(p)
	}
}
