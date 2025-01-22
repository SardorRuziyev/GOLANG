package main

import (
	"fmt"
)

func main() {

	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi, 0, 1, 2, 3, 4, 4, 5, 6, 6)
	fmt.Println(xi)

}
