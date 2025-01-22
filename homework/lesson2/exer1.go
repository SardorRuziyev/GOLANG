package main

import "fmt"

func main() {
	x := [5]int{}
	for i := 0; i < 5; i++ {
		x[i] = i
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}

}

/*
Using a COMPOSITE LITERAL:
● create an ARRAY which holds 5 VALUES of TYPE int
● assign VALUES to each index position.
● Range over the array and print the values out.
○ print out the VALUE and the TYPE
*/
