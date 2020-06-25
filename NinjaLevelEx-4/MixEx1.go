package main

import (
	"fmt"
)

func main() {
	js := []string{"James", "Bond", "Shaken, not stirred"}
	ms := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(js)
	fmt.Println(ms)
	ss := [][]string{js, ms} //a slice of string ([][]string)
	fmt.Println(ss)
	fmt.Println("This is Range over the record")
	for i, v := range ss {
		fmt.Println("Range over the Record:", i, v)
	}
	//This is Range over the each data

	for j, k := range ss {
		fmt.Println("Record:", j)
		for j, m := range k {
			fmt.Printf("\t index position: %v \t value: \t%v\n", j, m)
		}
	}

}
