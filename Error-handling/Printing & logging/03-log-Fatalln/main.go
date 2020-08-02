package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("Name.txt")
	if err != nil {
		//fmt.Println("Error happened: ", err)
		//log.Println("Error happened: ", err)
		log.Fatalln("Error happened: ", err)
	}
}
//2020/08/02 16:47:37 Error happened:  open Name.txt: The system cannot find the file specified.
//exit status 1
