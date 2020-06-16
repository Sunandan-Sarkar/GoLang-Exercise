package main

import (
	"fmt"
)

func main() { //for condition { }
	i := 8 //"initialization" can be before for
	for i <= 10 {
		fmt.Println(i)
		i++ //"post" can be after println
	}

}
