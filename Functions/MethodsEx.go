package main

import (
	"fmt"
)

type Account struct {
	Name   string
	Number int
	Date   []string
	Age    int
}
type PremiumAccount struct {
	Account
	SpecialService bool
	Valid          bool
}

func main() {
	pa1 := PremiumAccount{
		Account: Account{Name: "Sunandan Sarkar",
			Number: 1234567890,
			Date:   []string{"June 25th,2020"},
			Age:    34,
		},
		SpecialService: true,
		Valid:          true,
	}

	pa2 := PremiumAccount{
		Account: Account{Name: "Novi Sarkar",
			Number: 9876543210,
			Date:   []string{"May 25th,2020"},
			Age:    28,
		},
		SpecialService: true,
		Valid:          true,
	}

	pa3 := Account{
		Name:   "Modon Sarkar",
		Number: 5678904321,
		Date:   []string{"December 25th,2020"},
		Age:    50,
	}
	fmt.Println(pa1)
	fmt.Println(pa2)
	fmt.Println(pa3)
	fmt.Println(pa1.Account.Name, pa2.Name, pa3.Name)
	fmt.Println(pa2.SpecialService, pa2.SpecialService)
	fmt.Println(pa1.Account.Number, pa2.Account.Number, pa3.Number)
	fmt.Println(pa3.Date, pa1.Date, pa2.Date)
	fmt.Println("------------Checkout the Methods here---------------")
	pa1.aco()
	pa2.aco()
	pa3.con()
}

//Creating methods here
func (pa PremiumAccount) aco() {
	fmt.Println("I am:", pa.Name)
	fmt.Println("My info:", pa.Age, pa.Number, pa.Date, pa.SpecialService, pa.Valid)
}
func (a Account) con() {
	fmt.Println("I am:", a.Name)
	fmt.Println("My info:", a.Age, a.Number, a.Date)
}
