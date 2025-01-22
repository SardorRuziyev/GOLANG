package main

import "fmt"

func main() {
	xs := []string{"sam", "cam", "dam", "ham", "pam", "ram"}
	fmt.Println(xs)
	fmt.Println("_______________")
	xs = append(xs, "lam", "kam", "yam", "wam")
	fmt.Println(xs)
	fmt.Println("_______________")

	fmt.Printf("xs - %v\n", xs[0:3])

}
