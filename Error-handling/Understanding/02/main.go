package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string

	fmt.Println("Your Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Favourite Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Favourite drink: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer1, answer2, answer3)

}
