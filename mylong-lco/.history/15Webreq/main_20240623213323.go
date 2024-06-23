package main

import (
	"fmt"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Types of response %T\n", response)

	// It is a caller responsiblity to close the connection
	defer response.Body.Close()
}
