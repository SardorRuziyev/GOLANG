package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	person := Person{
		FirstName: "ALI",
		LastName:  "VALIEV",
	}
	fmt.Println("full Name: ", person.FullName())

}
