package main

import "fmt"

type students struct {
	firstName string
	lastName  string
	age       int
}
type major struct {
	class string
	club  string
}

func main() {
	s1 := students{
		firstName: "Akmal",
		lastName:  "Umrzaqov",
		age:       38,
	}
	s2 := students{
		firstName: "Farhod",
		lastName:  "Ahunjanov",
		age:       45,
	}

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("%T \t %#v\n", s1, s1)

	fmt.Println(s1.firstName, s1.lastName, s1.age)
}
