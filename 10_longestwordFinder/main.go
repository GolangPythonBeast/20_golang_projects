package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	wordArray := wordLength()

	longestWord := ""
	maxlength := 0

	for _, word := range wordArray { //How to use range in golang
		if len(word) > maxlength {
			longestWord = word
			maxlength = len(word)
		}
	}
	fmt.Printf("The longest word is %s", longestWord)
}

func wordLength() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sentence or paragraph: ")
	text, _ := reader.ReadString('\n')

	wordArray := strings.Fields(text)
	return wordArray
}
