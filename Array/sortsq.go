package main

import (
	"fmt"
	"sort"
)

func sortSquares(slc []int) {

	for i := range slc {
		slc[i] = slc[i] * slc[i]
	}

	// Sort the squared elements
	sort.Ints(slc)
}

func main() {
	slc := []int{-4, -1, 0, 3, 10}

	fmt.Println("Before sort:")
	for _, val := range slc {
		fmt.Print(val, " ")
	}

	sortSquares(slc)

	fmt.Println("\nAfter sort:")
	for _, val := range slc {
		fmt.Print(val, " ")
	}
}
