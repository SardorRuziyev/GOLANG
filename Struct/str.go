package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}
type Library struct {
	Name  string
	Books []Book
}

func (b Book) summary() string {
	return fmt.Sprintf("%s by %s (%d)", b.Title, b.Author, b.Year)
}

func main() {
	// book := Book{
	// 	Title:  "The Great Gatsby",
	// 	Author: "F. Scott Fitzgerald",
	// 	Year:   1925,
	// }
	var book Book
	book.Title = "The Great Gatsby"
	book.Author = "F. Scott Fitzgerald"
	book.Year = 1925

	fmt.Println(book)
	fmt.Println("______________")

	title := book.Title
	fmt.Println(title)
	fmt.Println("______________")

	library := Library{
		Name: "City Library",
		Books: []Book{
			{
				Title:  "1984",
				Author: "George Orwell",
				Year:   1949,
			},
			{
				Title:  "To Kill a Mockingbird",
				Author: "Harper Lee",
				Year:   1960,
			},
		},
	}

	fmt.Println(library)
	fmt.Println("______________")

	fmt.Println(book.summary())

}
