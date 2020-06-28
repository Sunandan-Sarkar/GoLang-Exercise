package main

import (
	"fmt"
)

func main() {
	x := []int{11, 33, 55, 77}
	y := []string{"Eleven", "Thirty three", "Fifty five", "Seventy seven"}
	z := []int{22, 44, 66, 88}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	ip := [][]int{x, z}

	fmt.Println("The value of ip:", ip)
	x = append(x, 99, 100)
	fmt.Println(x)

	x = append(x, z...)
	fmt.Println(x)

	fmt.Println("Deleting-------")
	x = append(x[2:5], x[8:]...)
	fmt.Println(x)

	y = append(y, "Eighty eight", "Ninenty Nine", "Hundred")
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))
	y = append(y[2:5])
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	m := map[string]int{
		"Digits": 4,
		"Number": 9,
	}
	fmt.Println(m)
	fmt.Println(m["Number"])
	fmt.Println("Deleting-------")
	delete(m, "Digits")
	fmt.Println(m)

	p := map[string]string{
		"Digits": "Four",
		"Number": "Nine",
	}
	fmt.Println(p)
	fmt.Println(p["age"])
	v, ok := p["age"]
	fmt.Println(v, ok)

	for i, v := range ip {
		fmt.Println("The record of ", i)
		for i, j := range v {
			fmt.Println("\t\t", i, "\t", j)
		}
	}

}
