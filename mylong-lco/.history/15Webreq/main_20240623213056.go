package main

import (
	"fmt"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)
}
