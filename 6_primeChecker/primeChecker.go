package main

import "fmt"

func main() {
	primeCheckToCheck := inputValue("Enter number to check: ")
	dataCheck := primeCheck(primeCheckToCheck)

	switch dataCheck {
	case true:
		fmt.Printf("%d is a prime number", primeCheckToCheck)
	case false:
		fmt.Printf("%d is not a prime number", primeCheckToCheck)
	}
}

func inputValue(text string) int {
	var numData int
	fmt.Print(text)
	_, err := fmt.Scan(&numData)
	if err != nil {
		fmt.Println("Invalid Input")
		return 0
	}
	return numData
}

func primeCheck(numb int) bool {
	var factors int = 0

	for i := 1; i <= numb; i++ {
		if numb%i == 0 {
			factors++
		}
	}
	if factors == 2 {
		return true
	} else {
		return false
	}
}
