package main

import "fmt"

type Student struct {
	ID   int
	Name string
	Age  int
}

type Class struct {
	Name     string
	Students []Student
}

func main() {
	student1 := Student{ID: 1, Name: "Sam", Age: 20}
	student2 := Student{ID: 2, Name: "Cam", Age: 20}
	student3 := Student{ID: 3, Name: "Liam", Age: 25}
	student4 := Student{ID: 4, Name: "Dan", Age: 24}
	student5 := Student{ID: 5, Name: "Sami", Age: 23}

	class1 := Class{
		Name:     "Mathematics",
		Students: []Student{student1, student2, student5},
	}
	class2 := Class{
		Name:     "English",
		Students: []Student{student2, student3, student4},
	}

	classes := []Class{class1, class2}

	for _, class := range classes {

		fmt.Printf("Class Name: %s\n", class.Name)

		fmt.Println("Students in the class:")
		for _, student := range class.Students {
			fmt.Printf("  ID: %d, Name: %s, Age: %d\n", student.ID, student.Name, student.Age)
		}
	}
}
