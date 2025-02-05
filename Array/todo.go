package main

import "fmt"

func main() {

	tasks := []string{"Buy groceries", "Clean room", "Call mom"}
	fmt.Println("to-do list:")
	for i := 0; i < len(tasks); i++ {
		fmt.Println("-", tasks[i])
	}
	fmt.Println("Updated to-do list")
	tasks = append(tasks[:1], tasks[2:]...)
	for i := 0; i < len(tasks); i++ {
		fmt.Println("-", tasks[i])
	}

}
