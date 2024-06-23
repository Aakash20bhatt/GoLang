package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006")) // Standard for the format mm-dd-yyyy

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(2020, time.August, 20, 23, 23, 0, 1, time.Now().Location())

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
