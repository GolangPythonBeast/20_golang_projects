package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Syntax: ./parser <json_file>")
		return
	}

	json_file := os.Args[1]                // How use use terminal argument in the code
	content, err := os.ReadFile(json_file) // How to read file using the os package

	if err != nil {
		fmt.Println("Error reading from file")
		return
	}

	var user User

	err = json.Unmarshal(content, &user) // How to read the json file using the json package

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Name: %s\n", user.Name)
		fmt.Printf("Age: %d", user.Age)
	}

}
