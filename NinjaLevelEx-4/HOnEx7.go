package main

import (
	"fmt"
)

func main() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	fmt.Println(jb)
	fmt.Println(mp)

	jp := [][]string{jb, mp}
	fmt.Println(jp)

	for i, val := range jp {
		fmt.Println("Record: ", i)
		for i, pos := range val {
			fmt.Printf("\t index position: %v \t value: \t %v \n", i, pos)
		}

	}

}
