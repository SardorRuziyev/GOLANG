package main

import "fmt"

func sumstring(s string) int {
	sum := 0

	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}
	return sum
}
func main() {
	s := "abc"
	sum := sumstring(s)
	fmt.Printf("The sum of the string '%s' is: %d\n", s, sum)
}
