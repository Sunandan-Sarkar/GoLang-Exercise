package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"John","Last":"Snow","Age":33},{"First":"White","Last":"Man","Age":20}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("All the data", people)
	for i, v := range people {
		fmt.Println("People number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
