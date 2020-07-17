package main

import (
	"fmt"
	"sort"
)

type us struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type ByAge []us

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []us

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

func main() {
	u1 := us{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := us{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := us{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	fmt.Println("------------------")
	users := []us{u1, u2, u3}
	for i,v:=range users {
		fmt.Println("People #", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t", v2)
		}
	}
	fmt.Println("------------------")
	sort.Sort(ByAge(users))
	for i,v:=range users{
		fmt.Println("\t People #",i)
		fmt.Println("\t\t",v.First,v.Last,v.Age)
		for _,v2:=range v.Sayings{
			fmt.Println("\t",v2)
		}
	}
	fmt.Println("------------------")
	sort.Sort(ByLast(users))
	for i,v:=range users{
		fmt.Println("People #", i)
		fmt.Println("\t",v.First,v.Last)
		for _,v2:=range v.Sayings{
			fmt.Println("\t",v2)
		}
	}

}

