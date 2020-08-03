package main

import "fmt"

func main() {
	//if false{
	//	fmt.Println("Print Tallinn")
	//} else if 4==3 {
	//	fmt.Println("Don't Print Tallinn")
	//} else if 3==4{
	//	fmt.Println("Tallinn")
	//} else {
	//	fmt.Println("Print Tallinn")
	//}
	x := "Limbaroyati"
	if x == "Sundan" {
		fmt.Println("Print Sundan")
	} else if x == "Limbaroyati" {
		fmt.Println("Print Limbaroyati")
	} else if x == "Oxyzen" {
		fmt.Println("Oxyzen")
	} else {
		fmt.Println("Default")
	}

}
