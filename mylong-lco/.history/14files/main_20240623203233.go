package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	content := "This needs to go in a file"

	file, err := os.Create("./mytextfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("length is: ", length)

	defer file.Close()

	readFile("./mytextfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	checkNilError(err)
	fmt.Println("Text data inside the file is :-", string(databyte))

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
