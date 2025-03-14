package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	fmt.Println("Profile Information: ")
	person1 := Person{
		Name:  "Bob",
		Age:   23,
		Email: "example1@gmail.com",
	}
	displayInformation(person1)
	person2 := Person{
		Name:  "Tunji",
		Age:   25,
		Email: "example2@gmail.com",
	}
	displayInformation(person2)

}

func displayInformation(p Person) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Email: %s\n", p.Email)
	fmt.Println()
}
