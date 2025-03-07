package main

import (
	"fmt"
	"math"
)

func main() {
	weight := addInput("Enter your weight (kg): ")
	height := addInput("Enter your height (m): ")
	heightPower := math.Pow(height, 2.0)
	bmi := weight / heightPower
	fmt.Printf("Your BMI is %.2f\n", bmi)

	bmiCal(bmi)
}

func addInput(text string) float64 {
	var dataEntered float64
	fmt.Print(text)
	fmt.Scan(&dataEntered)
	return dataEntered

}

func bmiCal(bmi float64) {
	// below is our to use control statement in Go

	if bmi >= 30 {
		fmt.Println("Category: Obese")
		return
	} else if bmi >= 25 && bmi <= 29.9 {
		fmt.Println("Category: Over weight")
		return
	} else if bmi >= 18.5 && bmi <= 24.9 {
		fmt.Println("Category: Normal weight")
		return
	} else {
		fmt.Println("Category: Underweight")
		return
	}
}
