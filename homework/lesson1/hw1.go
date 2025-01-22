package main

import (
	"fmt"
	"time"
)

func main() {

	// Displaying the full name
	firstName := "Sardor"
	lastName := "Ruziyev"

	fmt.Println(firstName + " " + lastName + "\n")

	//  Displaying the age
	birthYear := 1986
	age := time.Now().Year() - birthYear
	fmt.Println("Your age is", age)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d squared is %d\n", i, i*i)
	}
}
