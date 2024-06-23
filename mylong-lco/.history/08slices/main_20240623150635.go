package main

import "fmt"

func main() {

	fmt.Println("Welcome to the slices")

	var fruitlist = []string{"April", "May", "June"}
	fmt.Println("Types of fruitList is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

}
