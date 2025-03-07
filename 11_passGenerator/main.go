package main

import (
	"fmt"
	"math/rand"
	"time"
)

const pool = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXy"

func main() {
	password_length := 10

	var password string

	rand.NewSource(time.Now().UnixNano())
	for i := 0; i <= password_length; i++ {
		character := pool[rand.Intn(len(pool))]
		password += string(character)
	}
	fmt.Println("Generated Password:", password)

}
