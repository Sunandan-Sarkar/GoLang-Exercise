package main

import "fmt"

func main() {
	m := map[string]int{
		"Limbaroyati": 28,
		"Sunandan":    34,
		"Novi":        50,
	}
	fmt.Println(m)

	fmt.Println(m["Novi"])
	fmt.Println(m["Sun"])

	v, ok := m["Sun"]
	fmt.Println(v)
	fmt.Println(ok)

	//add a new element to map
	m["Tim"] = 69
	m["Covid"] = 19
	fmt.Println(m)

	fmt.Println("The below one is if value")

	if v, ok := m["Novi"]; ok {
		fmt.Println("Then print IF value", v)
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	xi := []int{911, 822, 733, 644, 555, 466, 377, 288, 199}
	for j, v := range xi {
		fmt.Println(j, v)
	}
	//Alternative way
	for i := 0; i <= 8; i++ {
		fmt.Println(i, xi[i])
	}

}
