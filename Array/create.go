package main

import "fmt"

type DVD struct {
	Name        string
	ReleaseYear int
	Director    string
}

func (d DVD) String() string {
	return fmt.Sprintf("%s, directed by %s, released in %d", d.Name, d.Director, d.ReleaseYear)
}

func main() {

	dvdCollection := make([]*DVD, 15)

	dvdCollection[0] = &DVD{
		Name:        "Inception",
		ReleaseYear: 2010,
		Director:    "Christopher Nolan",
	}

	dvdCollection[7] = &DVD{
		Name:        "The Avengers",
		ReleaseYear: 2012,
		Director:    "Joss Whedon",
	}

	fmt.Println("Complete DVD Collection:")
	for i, dvd := range dvdCollection {
		if dvd != nil {
			fmt.Printf("Index %d: %s\n", i, dvd)
		} else {
			fmt.Printf("Index %d: Empty\n", i)
		}
	}

	fmt.Println("Length of the Array: ", len(dvdCollection))
	fmt.Println("Capacity of the Array: ", cap(dvdCollection))

}
