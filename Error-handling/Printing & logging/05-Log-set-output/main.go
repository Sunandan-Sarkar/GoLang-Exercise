package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	f,err:=os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2,err:=os.Open("Name.txt")
	if err != nil {
		//fmt.Println("Error happened",err)
		//log.Println("Error happened",err)
		//log.Fatalln("Error happened",err)
		panic(err)
	}
	defer f2.Close()
	fmt.Println("Check the log.txt file in the directory")
}
//Error happened open Name.txt: The system cannot find the file specified.
//Check the log.txt file in the directory

//Check the log.txt file in the directory

//exit status 1

/*panic: open Name.txt: The system cannot find the file specified.

goroutine 1 [running]:
main.main()
        C:/Users/sun90/Documents/goworkspace/src/github.com/Sunandan-Sarkar/GoLang-Exercise/Error-handling/Printing & logging/05-Log-set-output/main.go:22 +0x213
exit status 2
*/
