package main

import "fmt"

func main() {
	isPrime := false
	for i := 1; i <= 30; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			} else {
				isPrime = true
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}

}
