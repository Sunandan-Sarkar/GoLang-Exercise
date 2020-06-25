package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(m)

	m[`Sun_dan`] = []string{`Being human`, `Burger`, `Movies`}
	fmt.Println(m)
	delete(m, `moneypenny_miss`)

	for i, v := range m {
		fmt.Println(i)
		for j, k := range v {
			fmt.Printf("\t index position: %v \t value: %v\n", j, k)
		}
	}
}
