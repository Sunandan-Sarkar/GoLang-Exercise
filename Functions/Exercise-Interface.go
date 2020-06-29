package main

import (
	"fmt"
)
type student struct {
	name string
	id int
	year int
	major string
}
type Sstudent struct {
	student
	completed bool
	remarks []string
}
func main(){
	s1:=student{
		name:  "Modon Lal",
		id:    1122334455,
		year:  2020,
		major: "Social Science",
	}
	s2:=Sstudent{
		student:student{
			name:  "Robin Hood",
			id:    6677889900,
			year:  2023,
			major: "Computer Engineering",
		},
		completed: false,
		remarks: []string{"Special","Helpful","Social"},
	}
	s3:=Sstudent{
		student:student{
			name:  "Moris Green",
			id:    7788554433,
			year:  2021,
			major: "Gaming VR",
		},
		completed: false,
		remarks: []string{"Special","Obedient","Honest"},
	}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	s1.soo()
	s2.soo()
	s3.soo()
	soo(s1)
	soo(s2)
	soo(s3)
}
func(ss1 student)soo(){
	fmt.Println(ss1.name,ss1.id,ss1.major,ss1.year)
}

type TLUStudent interface {
	soo()
}
func soo(s TLUStudent){
	switch s.(type) {
	case student:
		fmt.Println("I am normal",s)
	case Sstudent:
		fmt.Println("I am special",s)
	}
	fmt.Println("I am",s)
}
