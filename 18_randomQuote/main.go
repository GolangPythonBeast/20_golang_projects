package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://zenquotes.io/api/random"

type QuoteContent struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var quotes []QuoteContent
	err = json.Unmarshal(result, &quotes)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(quotes) > 0 {
		fmt.Printf("'%s' - %s\n", quotes[0].Q, quotes[0].A)
	} else {
		fmt.Println("No quote found in response.")
	}

}
