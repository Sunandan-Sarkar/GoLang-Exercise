package main

import "fmt"

func main() {
	bd := []string{"Sun", "Moon", "Star", "Jupitar", "Earth", "Planet"}
	fmt.Println(bd)
	bd[5] = "Earth"
	fmt.Println(bd)
	fmt.Println(len(bd))
	fmt.Println(bd[5])
	fmt.Println(bd[2:5])

	fmt.Println("Here is For loop")

	for i, v := range bd {
		fmt.Println(i, v)

	}
	fmt.Println("The below one is Alternative method")
	for i := 0; i <= 5; i++ {
		fmt.Println(i, bd[i])
	}

}
