package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("-----------")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		merr := fmt.Errorf("My error is %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", merr}
	}
	return 42, nil
}
