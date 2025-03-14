package main

import (
	"fmt"
)

type Contact struct {
	Name        string
	PhoneNumber int
}

var contactList []Contact

func addContact() {
	var name string
	fmt.Print("Enter the contact name: ")
	fmt.Scan(&name)

	var phoneNumber int
	fmt.Print("Enter the phone number: ")
	fmt.Scan(&phoneNumber)

	contact := Contact{
		Name:        name,
		PhoneNumber: phoneNumber,
	}

	// creating content for the struct
	contactList = append(contactList, contact)
	fmt.Println("Contact Added Successfully!")
}

func listContacts() {
	if len(contactList) == 0 {
		fmt.Println("No saved contact")
		fmt.Println()
	}
	for index, contact := range contactList {
		fmt.Printf("%d. %s - %d", index+1, contact.Name, contact.PhoneNumber)
	}
	fmt.Println()
}

func main() {
	fmt.Println("Contact List App")

	for {

		fmt.Println("\n1. Add contact")
		fmt.Println("2. List contacts")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addContact()
		case 2:
			listContacts()
		case 3:
			fmt.Println("Exiting.......")
			return
		default:
			fmt.Println("Invalid Input")
		}
	}
}
