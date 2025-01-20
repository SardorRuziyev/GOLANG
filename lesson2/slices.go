package main

import "fmt"

func main() {
	xs := []string{"sam", "cam", "dam", "ham", "pam", "ram"}
	fmt.Println(xs)
	fmt.Printf("%T\n", xs)

	for i, v := range xs {
		fmt.Printf(" index %v - value %v\n ", i, v)
	}

	for _, v := range xs {
		fmt.Printf(" %v\n ", v)
	}

	fmt.Println(len(xs))
}
