package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web verb video")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	fmt.Println("Content length", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
