package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}

type student struct {
	person
}

type teacher struct {
	person
	major string
	group group
}

type group struct {
	name     string
	class    int
	students []student
}

func main() {
	student3 := student{person{name: "Jamshid", lastname: "Bola", age: 23}}
	fmt.Println(student3.name)
	group3 := group{students: []student{student3}}
	group3.name = "English"
	group3.class = 765
	fmt.Println(group3)
	t := teacher{}
	t.name = "Jon"
	t.major = "History"
	t.group.class = 6543
}
