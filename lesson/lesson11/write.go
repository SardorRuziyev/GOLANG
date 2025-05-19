package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	username, orderNumber := "Adams_adebayo", "ORD6543234"
	file, err := os.Create(username + orderNumber + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	item1, item2, item3 := "shoe", "bag", "shirt"
	price1, price2, price3 := 1000, 2000, 3000

	_, err = fmt.Fprintf(file, "Username: %s\nOrder Number: %s\nItem 1: %s\nPrice 1: %d\nItem 2: %s\nPrice 2: %d\nItem 3: %s\nPrice 3: %d\n",
		username, orderNumber, item1, price1, item2, price2, item3, price3)
	if err != nil {
		log.Fatal(err)
	}
}
