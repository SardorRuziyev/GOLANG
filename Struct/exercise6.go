package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float64
}

type University struct {
	Name     string
	Students []Student
}

func (u *University) BestStudent() Student {
	best := 

	for _, student := range u.Students {
		if student.GPA > best.GPA {
			best = student
		}
	}

	return best
}

func main() {
	students := []Student{
		{ID: 1, Name: "John", GPA: 3.9},
		{ID: 2, Name: "Ben", GPA: 3.1},
		{ID: 3, Name: "Ken", GPA: 4.0},
	}

	university := University{"TATU", students}
	fmt.Println(university.BestStudent())
}
