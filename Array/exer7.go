package main

import "fmt"

func main() {
	numbers := []int{10, 15, 21, 7, 30, 13}

	for i := 0; i < len(numbers); i++ {
		isPrime := false

		for j := 2; j < i; j++ {
			if i % j == 0 {
				isPrime = false
				break
			} else {
				isPrime = true
			}
		}
		if isPrime {
			fmt.Print(numbers[i], " ")
			break
		}

	}

}
