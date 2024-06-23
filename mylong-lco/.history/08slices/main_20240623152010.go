package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to the slices")

	var fruitlist = []string{"April", "May", "June"}
	fmt.Println("Types of fruitList is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highScore := make([]int, 4)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867

	fmt.Println(highScore) // If we add another value then we get error eg. highScore[4]=777

	highScore = append(highScore, 555, 666, 321) // Assign new memory

	sort.Ints(highScore)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "js", "swift", "ruby", "swift"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
