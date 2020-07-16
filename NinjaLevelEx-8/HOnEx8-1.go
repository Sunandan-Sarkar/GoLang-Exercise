package main

import (
	"encoding/json"
	"fmt"
)

type usr struct {
	First string
	Age   int
}

func main() {
	u1 := usr{
		First: "James",
		Age:   32,
	}

	u2 := usr{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := usr{
		First: "M",
		Age:   54,
	}

	users := []usr{u1, u2, u3}

	fmt.Println(users)

	us, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(us))
}

