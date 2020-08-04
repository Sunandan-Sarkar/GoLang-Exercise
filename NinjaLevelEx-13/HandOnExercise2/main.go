package main

import (
	"fmt"
	"github.com/Sunandan-Sarkar/GoLang-Exercise/NinjaLevelEx-13/HandOnExercise2/quote"
	"github.com/Sunandan-Sarkar/GoLang-Exercise/NinjaLevelEx-13/HandOnExercise2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}


