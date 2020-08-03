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
		//log.Fatalln("Error happened: ", err)
		//panic(err)
		log.Panic(err)
	}
}

//panic: open Name.txt: The system cannot find the file specified.
//
//goroutine 1 [running]:
//main.main()
//        C:/Users/sun90/Documents/goworkspace/src/github.com/Sunandan-Sarkar/GoLang-Exercise/Error-handling/Printing & logging/04-panic/main.go:13 +0x7b
//exit status 2
