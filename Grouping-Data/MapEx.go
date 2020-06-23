package main

import (
	"fmt"
)

func main() {
	dy := map[string]int{
		"Sunday":    9,
		"Monday":    10,
		"Tuesday":   11,
		"Wednesday": 12,
		"Thursday":  13,
		"Friday":    14,
		"Saturday":  15,
		"Mayday":    16,
	}
	fmt.Println(dy)
	fmt.Println(dy["Sunday"])
	fmt.Println(dy["Holyday"])

	v, ok := dy["Holyday"]
	fmt.Println(v, ok)

	if v, ok := dy["Sunday"]; ok {
		fmt.Println(v, ok)
	}
	dy["Offday"] = 24 // adding new element
	fmt.Println(dy)

	for d, v := range dy {
		fmt.Println(d, v)
	}
	di := []int{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24}
	for i, v := range di {
		fmt.Println(i, v)
	}
	fmt.Println("The below one is alternative method")
	for j := 0; j <= 15; j++ {
		fmt.Println(j, di[j])
	}
}
