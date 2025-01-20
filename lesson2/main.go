package main

import "fmt"

func main() {
	// array literal
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "Students!", "How are you!"}
	fmt.Println(b)

	var c [2]int
	c[0] = 3
	c[1] = 4
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
