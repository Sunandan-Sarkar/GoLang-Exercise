package main

import (
	"fmt"
)

type mobile struct {
	model   string
	brand   string
	release int
	rating  int
}
type SEmobile struct {
	mobile
	Sfunction bool
	price []string
}

func main() {
	p1 := mobile{
		model:   "iPhone5",
		brand:   "Apple",
		release: 2012,
		rating:  9,
	}
	p2 := SEmobile{
		mobile: mobile{
			model:   "Nokia 6",
			brand:   "Nokia",
			release: 2016,
			rating:  6,
		},
		Sfunction: false,
		price: []string{"400","USD","Retail Price"},
	}
	p3:=SEmobile{
		mobile:    mobile{
			model:   "iPhoneX",
			brand:   "Apple",
			release: 2018,
			rating: 10,
		},
		Sfunction: true,
		price:     []string{"1500","EURO","Retail Price, First come first serve"},
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	defer p2.aoo()
	p1.boo()
	p3.coo()
}
func (sp1 SEmobile) aoo()  {
	fmt.Println(sp1.model)
	fmt.Println(sp1.brand,sp1.rating,sp1.release,sp1.Sfunction)
	fmt.Println(sp1.price)
}
func (sp2 mobile) boo(){
	fmt.Println(sp2.model)
	fmt.Println(sp2.brand,sp2.release,sp2.rating)
}
func(sp3 SEmobile)coo(){
	fmt.Println(sp3.brand,sp3.model,sp3.release,sp3.rating,sp3.Sfunction)
	fmt.Println(sp3.price)
}
