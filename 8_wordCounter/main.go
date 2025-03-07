package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a sentence or paragraph: ")

	text, _ = reader.ReadString('\n') // this helps to get multiple strings from the terminal

	wordCount := strings.Fields(text)
	data := len(wordCount)
	fmt.Printf("There are %d words altogether", data)

}
