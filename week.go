package main

import (
	"fmt"
)

func main() {
	var day string = ""
	var dayWords = ""
	var option string = "y"

	for option == "y" {

		fmt.Println("Please enter a number between 1 and 7")

		fmt.Scanln(&day)

		switch day {
		case "1":
			dayWords = "Monday"
		case "2":
			dayWords = "Tuesday"
		case "3":
			dayWords = "Wednesday"
		case "4":
			dayWords = "Thursday"
		case "5":
			dayWords = "Friday"
		case "6":
			dayWords = "Saturday"
		case "7":
			dayWords = "Sunday"
		}

		fmt.Println("The date is " + dayWords)
		fmt.Println("Do you want other day(y/n)")
		fmt.Scanln(&option)
	}
}
