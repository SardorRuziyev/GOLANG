package main

import "fmt"

func main() {
	var mehmonlar []string = []string{"ali", "vali", "hasan"}

	for i := 0; i < len(mehmonlar); i++ {
		fmt.Println("Salom, ", mehmonlar[i])
	}
}
