package main

import "fmt"

func main() {
	fmt.Println("Shopping List: ")
	shopping := []string{"Milk", "Bread", "Egg", "Butter"}
	for i := 0; i < len(shopping); i++ {
		fmt.Println("-", shopping[i])
	}
	shopping = append(shopping, "Banana")
	fmt.Println("-", shopping[len(shopping)-1])

}
