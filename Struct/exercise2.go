package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	book := Book{
		Title:  "Common Ground",
		Author: "Hamza Yusuf",
		Year:   2004,
	}
	fmt.Println("Book Detail: ")
	fmt.Println(book.Title, book.Author, book.Year)
}
