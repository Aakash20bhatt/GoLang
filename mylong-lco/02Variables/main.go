package main

import "fmt"

const LoginToken string = "ghabbhhjd" // This is actually a public
// In go language, if firstLetter is in capital, it means that variable is public

func main(){
	var username string = "Aakash";
	fmt.Println(username);
	fmt.Printf("Variable is of type: %T \n",username);

	var isloggedIn bool = false;
	fmt.Println(isloggedIn);
	fmt.Printf("Variable is of type: %T \n",isloggedIn);

	var smallVal uint8 = 255;
	fmt.Println(smallVal);
	fmt.Printf("Variable is of type: %T \n",smallVal);

	var smallFloat float32 = 255.44552123;
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n",smallFloat);

	// default values and some aliases

	var anotherVariables int 
	fmt.Println(anotherVariables);
	fmt.Printf("Variables is of type: %T \n", anotherVariables)

	var anotherVariablesdtr string 
	fmt.Println(anotherVariablesdtr);
	fmt.Printf("Variables is of type: %T \n", anotherVariablesdtr);

	var anotherVariablesbool bool 
	fmt.Println(anotherVariablesbool);
	fmt.Printf("Variables is of type: %T \n", anotherVariablesbool);

	// implicit type 

	var website = "aakash.com";
	fmt.Println(website);

	// no var style

	numberOfUser := 300000;
	fmt.Println(numberOfUser);
}