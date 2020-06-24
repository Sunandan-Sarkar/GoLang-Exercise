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

	m[`sun_dr`] = []string{`sunandan`, `Limbaroyati`, `sundan`}
	fmt.Println(m)

	for i, j := range m {
		fmt.Println("The record of", i)
		for p, k := range j {
			fmt.Println("\t", p, k)
		}
	}
}
