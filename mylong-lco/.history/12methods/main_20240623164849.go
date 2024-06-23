package main

import "fmt"

func main() {
	// No inheritance in golang, No super or parent
	aakash := User{"Aakash", "aakash@go.dev", true, 16}

	fmt.Println(aakash)
	fmt.Printf("User details are %+v\n", aakash) // %+v gives you more details
	fmt.Printf("Name is %v and email is %v:", aakash.Name, aakash.Email)
	aakash.GetStatus()
	aakash.NewMail()
	fmt.Printf("Name is %v and email is %v\n", aakash.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Its user active", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.com"
	fmt.Println("Email of this user is", u.Email)
}
