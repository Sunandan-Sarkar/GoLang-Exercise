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

	for i, j := range m {
		fmt.Println("This is the record of", i)
		for k, l := range j {
			fmt.Println("\t", k, l)
		}
	}
}
