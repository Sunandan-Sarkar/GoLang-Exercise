package main

import (
	"fmt"
)

func main() {
	x := []string{"Sun", "Moon", "Star", "Jupitar", "Earth", "Planet"}
	fmt.Println(x)
	x = make([]string, 10, 10) //make([]T, length, capacity)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = "Apples"
	x[1] = "Banana"
	x[2] = "Canberry"
	x[3] = "Dalim"
	x[4] = "Eggplant"
	x[5] = "Mango"
	x[6] = "Melon"
	x[7] = "Jackfruits"
	x[8] = "Avocados"
	x[9] = "Pear"
	fmt.Println(x)

	x = append(x, "Pineapple")

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, "Grapes")

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, "Strawberry")

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, "Orange", "Or", "Oge")

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
