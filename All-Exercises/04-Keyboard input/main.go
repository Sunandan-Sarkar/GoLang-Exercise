package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter you name: ")
	name,_:=reader.ReadString('\n')
	fmt.Print("Hello ",name)
}
