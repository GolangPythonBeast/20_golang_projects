package main

import "fmt"

func main() {
	bill := inputValue("Enter the bill ($) value: ")
	tipPercent := inputValue("Enter the tip (%) value: ")

	tip := 0.01 * tipPercent * bill
	totalBill := bill + tip

	fmt.Printf("Tip Amount: $%.2f\n", tip)
	fmt.Printf("Total Bill (including tip): $%.2f\n", totalBill)
}

func inputValue(text string) float64 {
	var inputData float64
	fmt.Print(text)
	fmt.Scan(&inputData)
	return inputData
}
