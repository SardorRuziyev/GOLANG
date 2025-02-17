package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}

func main() {
	user := User{
		ID:    1,
		Name:  "Eshmirza",
		Age:   23,
		Email: "eshmirza@gmail.com",
	}
	fmt.Println("User Details: ")
	fmt.Printf("ID:%d\n", user.ID)

}
