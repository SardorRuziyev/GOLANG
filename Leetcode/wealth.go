package main

import "fmt"

func main() {
	accounts := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{5, 6, 7},
	}
	// fmt.Println(accounts, accounts[0])
	fmt.Println(maximumWealth(accounts))
}
func maximumWealth(accounts [][]int) int {
	sum := 0

	for i := 0; i < len(accounts); i++ {
		currentsum := 0
		for j := 0; j < len(accounts[i]); j++ {
			currentsum += accounts[i][j]
		}
		if currentsum > sum {
			sum = currentsum
		}
	}

	return sum
}
