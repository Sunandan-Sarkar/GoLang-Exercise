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
	delete(dy, "Sunday")
	fmt.Println(dy)
	delete(dy, "Holyday")
	fmt.Println(dy)
	if v, ok := dy["Wednesday"]; ok {
		fmt.Println(v, ok)
		delete(dy, "Wednesday")
		fmt.Println(dy)
		delete(dy, "Friday")
		fmt.Println(dy)
		delete(dy, "Saturday")
		fmt.Println(dy)
		delete(dy, "Sunday")
		fmt.Println(dy)
		delete(dy, "Tuesday")
		fmt.Println(dy)
		delete(dy, "Mayday")
		fmt.Println(dy)
		delete(dy, "Monday")
		fmt.Println(dy)
		delete(dy, "Thursday")
		fmt.Println(dy)
	}

}
