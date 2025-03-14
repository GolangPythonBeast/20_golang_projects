package main

// This project is to make using of external packages easy
import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

func main() {
	fmt.Printf("Name: %s\n", faker.Name())
	fmt.Printf("EMail: %s\n", faker.Email())
	fmt.Printf("Username: %s\n", faker.Username())
	fmt.Printf("Password: %s\n", faker.Password())
}
