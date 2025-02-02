package main

import "fmt"

func main() {
	n := 5
	fac := 1

	for i := 1; i <= n; i++ {
		fac = fac * i
		// 1, 2, 6, 24, 120
		fmt.Print(fac, ", ")

	}
}
