package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "This needs to go in a file"

	file, err := os.Create("./mytextfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("length is: ", length)

	defer file.Close()
}
