package main

import "fmt"

func main() {
	temperatures := []float64{32, 31.5, 30.8, 42}

	sum := 0

	for i := 0; i < len(temperatures); i++ {
		sum += int(temperatures[i])
	}
	average := sum / len(temperatures)

	fmt.Println(average)

}
