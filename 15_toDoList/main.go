package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todos []string

func addToDo() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the task description: ")
	task, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	task = strings.TrimSpace(task) // Remove trailing newline
	todos = append(todos, task)
	fmt.Println("Task added successfully!")
}

func listToDo() {
	if len(todos) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("To-Do List 👇")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i+1, todo)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("To-Do List App")

	for {
		fmt.Println("1. Add To-Do item")
		fmt.Println("2. List To-Dos")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		// Use bufio to read input instead of fmt.Scan()
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addToDo()
		case "2":
			listToDo()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid Input! Try again.")
		}
	}
}
