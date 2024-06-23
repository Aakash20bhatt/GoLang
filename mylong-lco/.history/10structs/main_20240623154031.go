package main

import "fmt"

func main() {
	// No inheritance in golang, No super or parent
	aakash := User{"Aakash", "aakash@go.dev", true, 16}

	fmt.Println(aakash)
	fmt.Printf("User details are %+v\n", aakash)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
