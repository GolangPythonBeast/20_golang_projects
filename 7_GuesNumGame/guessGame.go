package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.NewSource(time.Now().UnixNano())
	target := rand.Intn(100) + 1 // this is a way to generate random number
	attempt := 0

	welcomeMessage()

	for attempt = 1; attempt <= 10; attempt++ {
		leftAttempts := 10 - attempt
		fmt.Print("Enter your guess: ")
		var guess int
		fmt.Scan(&guess)

		if target > guess {
			fmt.Printf("Too low! Try again (%d attempts remaining)\n", leftAttempts)
		} else if target < guess {
			fmt.Printf("Too high! Try again (%d attempts remaining)\n", leftAttempts)
		} else {
			fmt.Printf("Correct! You guessed right after %d attempts", attempt)
			break
		}
		if attempt == 10 {
			fmt.Println("Game Over")
			break
		}
	}

}

func welcomeMessage() {
	fmt.Print(`
Welcome to the Guess Number Game. 
I am thinking of a number between 1 and 100. Can you guess it?
`)

}
