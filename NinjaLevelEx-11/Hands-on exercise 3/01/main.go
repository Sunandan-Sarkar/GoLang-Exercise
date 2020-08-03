package main

import "fmt"

type customErr struct {
	first   string
	last    string
	sayings []string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is my error %v\n",ce.sayings)
}

func main() {

	c1 := customErr{
		first:   "MadMax",
		last:    "Funny",
		sayings: []string{"Life is beautiful", "Go in Action", "Life isn't bed of roses"},
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
