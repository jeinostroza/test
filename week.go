package main

import (
	"fmt"
)

func main() {
	var day string = ""
	var dayWords = ""

	fmt.Println("Please enter a number between 1 and 7")

	fmt.Scanln(&day)

	switch day {
	case "1":
		dayWords = "Monday"
	case "2":
		dayWords = "Tuesda"
	case "3":
		dayWords = "Wednesday"
	case "4":
		dayWords = "Thursday"
	}

	fmt.Println("The day is " + dayWords)
}
