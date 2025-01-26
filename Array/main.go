package main

import "fmt"

func main() {
	n := 20
	i := 20

	for i = 0; i < n; i-- {
		fmt.Println("--")
		break

	}
}

/*
Initialization (i := 0): Sets up the starting condition.
Condition (i < n): Runs the loop as long as the condition is true.
Post-statement (i++): Executes after each loop iteration.
*/
