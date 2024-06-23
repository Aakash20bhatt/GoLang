package main

import "fmt"

func main() {

	programmingLanguage := make(map[string]string)

	programmingLanguage["JS"] = "JavaScript"
	programmingLanguage["RB"] = "Ruby"
	programmingLanguage["Py"] = "Python"

	fmt.Println(programmingLanguage)

	delete(programmingLanguage, "RB")

	// loops in maps

	for key, value := range programmingLanguage {
		fmt.Println("For key %v, value is %v\n")
	}
}
