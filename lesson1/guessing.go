package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(101)

	fmt.Println("1 bilan 100 atrofidagi tahmin sonni topish")

	var guess int = -1
	for guess != number {

		fmt.Print("\nBirorta tahmin son o'ylang: ")
		fmt.Scan(&guess)

		if guess == number {
			fmt.Printf("Uraa, tahmin sonni topdingiz %d\n", number)
		} else if guess > number {
			fmt.Println("Siz tahmin qilgan son ancha baland")
		} else {
			fmt.Println("Tahminingiz ozgina pastroq")
		}
	}
}
