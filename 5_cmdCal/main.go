package main

import "fmt"

func main() {
	var firstNumber, secondNumber float64
	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&firstNumber)
	if err != nil {
		fmt.Print("Invalid Input")
		return
	}

	fmt.Print("Operator (+, -, *, /): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Print("Invalid Input")
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scan(&secondNumber)
	if err != nil {
		fmt.Print("Invalid Input")
		return
	}

	switch operator {
	case "+":
		result := (firstNumber + secondNumber)
		fmt.Printf("%.2f %v %.2f = %.2f", firstNumber, operator, secondNumber, result)
	case "-":
		result := (firstNumber - secondNumber)
		fmt.Printf("%.2f %v %.2f = %.2f", firstNumber, operator, secondNumber, result)
	case "*":
		result := (firstNumber * secondNumber)
		fmt.Printf("%.2f %v %.2f = %.2f", firstNumber, operator, secondNumber, result)
	case "/":
		if secondNumber == 0 {
			fmt.Println("Error: You can't divide with 0")
			return
		}
		result := (firstNumber / secondNumber)
		fmt.Printf("%.2f %v %.2f = %.2f", firstNumber, operator, secondNumber, result)
	default:
		strng := fmt.Sprintf("Invalid Operator: %v", operator)
		fmt.Print(strng)
	}

}
