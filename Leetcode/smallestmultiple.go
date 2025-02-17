package main

import "fmt"

func smallestMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}
func main() {

	n := 7
	fmt.Println(smallestMultiple(n))
}
