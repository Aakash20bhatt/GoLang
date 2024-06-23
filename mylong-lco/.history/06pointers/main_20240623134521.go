package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of the pointers ", ptr)
	fmt.Println("Value of the pointers ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("New value is ", myNumber)
}
