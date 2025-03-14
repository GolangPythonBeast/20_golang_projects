package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RandomJoke struct {
	Value string `json:"value"`
}

const url = "https://api.chucknorris.io/jokes/random"

func main() {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var joke RandomJoke
	err = json.Unmarshal(content, &joke)
	if err != nil {
		fmt.Sprintln(err)
		return
	}
	if joke == (RandomJoke{}) {
		fmt.Println("No joke return")
		return
	} else {
		fmt.Printf("%s\n", joke.Value)
	}

}
