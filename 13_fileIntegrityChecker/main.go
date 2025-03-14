package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

// 7509e5bda0c762d2bac7f90d758b5b2263fa01ccbc542ab5e3df163be08e6ca9
func main() {
	var inputValue int
	fmt.Print("Enter 1 to check original Hash or 2 to check integrity: ")
	fmt.Scan(&inputValue)

	if inputValue == 1 {
		var filename string
		fmt.Print("Enter the name of the file to hash: ")
		fmt.Scan(&filename)

		originalHash := getOriginalHash(string(filename))
		fmt.Println("The original string of the file is: ", originalHash)
	} else if inputValue == 2 {

		if len(os.Args) != 3 {
			fmt.Println("Usage: ./main.go <filename> <originalHash>")
			return
		}

		fileName := os.Args[1]
		originalHash := os.Args[2]

		contents, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
		}

		hash := hashedInput(contents)
		if hash == originalHash {
			fmt.Println("The file is intact")
		} else {
			fmt.Println("The file is not intact")
		}
	}

}

func getOriginalHash(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	hashValue := hashedInput(content)

	return hashValue
}

// The function gets hashed value of contents in a file
func hashedInput(content []byte) string {
	hasher := sha256.New()
	hasher.Write(content)
	result := hasher.Sum(nil)
	hash := hex.EncodeToString(result)
	return hash

}
