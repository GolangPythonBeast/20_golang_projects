package main

import "fmt"

func main() {
	var word string
	fmt.Print("Enter the word to check: ")
	fmt.Scan(&word)

	if word == reverseWord(word) {
		fmt.Printf("%s is a palindrome", word)
	} else {
		fmt.Printf("%s is not palindrome", word)
	}
}

func reverseWord(text string) string {
	reversed := ""
	sizeOfText := len(text) - 1
	for i := sizeOfText; i >= 0; i-- {
		reversed += string(text[i])
	}
	return reversed
}
