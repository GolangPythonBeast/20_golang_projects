package main

import (
	"fmt"
	"time"
)

//declearing packages

func main() {
	currentYear := time.Now().Year() // This helps to get the current year
	birthYear := askAge("Enter your birth year: ")
	age := currentYear - birthYear
	if birthYear > currentYear {
		fmt.Println("Invalid Input")
		return
	} else if age > 200 {
		fmt.Println("You don't exist")
		return
	}

	fmt.Printf("You are %d years old now", age)

}

func askAge(text string) int {
	var data int
	fmt.Print(text)
	fmt.Scanln(&data)
	return data

}
