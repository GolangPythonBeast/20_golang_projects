package main

import "fmt"

func main() {
	inputValue, err := inputParams()
	if err != nil {
		fmt.Println(err)
	}

	// if inputValue == 1 {
	// 	tempValue := tempInput()
	// 	fahValue := (tempValue * 9 / 5) + 32
	// 	fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", tempValue, fahValue)
	// } else if inputValue == 2 {
	// 	tempValue := tempInput()
	// 	celValue := (tempValue - 32) * 5 / 9
	// 	fmt.Printf("%.2f Fahrenheir is %.2f Celsius\n", tempValue, celValue)
	// }

	// Instead of if statement, we can use the switch statement

	switch inputValue {
	case 1:
		tempValue := tempInput()
		fahValue := (tempValue * 9 / 5) + 32
		fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", tempValue, fahValue)
	case 2:
		tempValue := tempInput()
		celValue := (tempValue - 32) * 5 / 9
		fmt.Printf("%.2f Fahrenheir is %.2f Celsius\n", tempValue, celValue)
	}
}

func inputParams() (int, error) {
	var dataInputted int
	fmt.Print(`
Temperature Converter. 
Choose and option:
1. Celsius to Fahrenheit
2. Fahrenheit to celsius

Enter your choice (1 or 2): `)
	_, err := fmt.Scan(&dataInputted) // this ensure users input is integer
	if err != nil {
		panic(err)
	}
	return dataInputted, nil
}

func tempInput() float64 {
	var tempValue float64
	fmt.Print("Enter the temperature: ")
	fmt.Scan(&tempValue)
	return tempValue
}
