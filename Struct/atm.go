package main

import "fmt"

type User struct {
	ID   int
	Name string
	Card Card
}

type Card struct {
	CardNumber string
	Balance    float64
	Pin        string
}

type ATM struct {
	Location      string
	CashAvailable float64
	Bank          Bank
}
type Bank struct {
	Name    string
	Accouts []Card
}

func (c *Card) CheckBalance() float64 {
	return c.Balance

}

func (w *Card) Withdraw(amount float64) {
	if amount <= w.Balance {
		w.Balance -= amount
	} else {
		fmt.Println("Insufficient Fund")
	}
}
func (d *Card) Deposit(amont float64) {
	d.Balance += amont

}
func (c *Card) ValidatePin(inputPin string) bool {
	return c.Pin == inputPin
}

func main() {
	bank := Bank{Name: "Hamkor"}

	card := Card{
		CardNumber: "379287106904788",
		Balance:    1000.0,
		Pin:        "1234",
	}

	bank.Accouts = append(bank.Accouts, card)

	user := User{
		ID:   1,
		Name: "Eshmirza",
		Card: card,
	}

	// atm := ATM{
	// 	Location:      "Andijon",
	// 	CashAvailable: 3000.0,
	// 	Bank:          bank,
	// }

	fmt.Printf("balance: ", user.Card.CheckBalance())
}
