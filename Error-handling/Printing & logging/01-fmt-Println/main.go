package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("Name.txt")
	if err != nil {
		fmt.Println("Error happened: ", err)
	}
}

//Error happened:  open Name.txt: The system cannot find the file specified.
