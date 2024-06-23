package main

import "fmt"

func main() {
	gretter()

	// Nested function in main
	// func greeterTwo(){
	// 	fmt.Println("Another Method")
	// }

	result := adder(3, 5)
	fmt.Println(result)

	proRes, myMessage := proAdder(2, 4, 5, 6)
	fmt.Println(proRes)
	fmt.Println(myMessage)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(value ...int) (int, string) {
	total := 0

	for _, val := range value {
		total += val
	}

	return total, "Hi this the total value"
}

func gretter() {
	fmt.Println("Namestey")
}

// func (){
// 	fmt.Println("Iffe")
// }()
