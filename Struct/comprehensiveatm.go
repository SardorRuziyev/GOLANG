package main

import (
	"fmt"
	"time"
)

type User struct {
	ID   int
	Name string
	Card *Card
}
type Card struct {
	CardNumber         string
	Balance            float64
	Pin                string
	TransactionHistory []Transaction
	FailedAttempts     int
	IsBlocked          bool
}
type Transaction struct {
	Amount    float64
	Type      string
	Timestamp time.Time
}
type ATM struct {
	Location      string
	CashAvailable float64
	Bank          *Bank
}
type Bank struct {
	Name     string
	Accounts map[string]*Card
}

func (c *Card) CheckBalance() float64 {
	return c.Balance
}
func (c *Card) Withdraw(amount float64, atm *ATM) bool {
	if c.IsBlocked {
		fmt.Println("card is bloacked")
		return false
	}
	if amount > c.Balance {
		fmt.Println("not enough")
		return false
	}
	if amount > atm.CashAvailable {
		fmt.Println("atm has no money")
		return false
	}

	c.Balance = c.Balance - amount

	atm.CashAvailable = atm.CashAvailable - amount

	c.TransactionHistory = append(c.Transaction, Transaction{
		Amount:    amount,
		Type:      "withdraw",
		Timestamp: time.Now(),
	})

	fmt.Println("new balance", c.Balance)

	return true

}

func (c *Card) Deposit(amount float64) {
	c.Balance = c.Balance + amount
	c.TransactionHistory = append(c.TransactionHistory, Transaction{
		Amount:    amount,
		Type:      "Deposit",
		Timestamp: time.Now(),
	})
	fmt.Println("yangi balans", c.Balance)
}

func (c *Card) ValidatePin(inputPin string) bool {
	if c.IsBlocked {
		fmt.Println("card is blocked")
		return false
	}
	if c.Pin == inputPin {
		c.FailedAttempts = 0
		return true
	}

	c.FailedAttempts++
	if c.FailedAttempts >= 3 {
		c.IsBlocked = true
		fmt.Println("contact with your bank")
	}
	fmt.Println("Wrong Pin")
	return false

}

func (c *Card) GetLastTransactions(n int) []Transaction {
	if n > len(c.TransactionHistory) {
		n = len(c.TransactionHistory)
	}
	return c.TransactionHistory[len(c.TransactionHistory)-n:]

}

func main() {
	bank := &Bank{
		Name:    "Hamkor",
		Accouts: make(map[string]*Card),
	}

	card1 := &Card{
		CardNumber:         "379287106904788",
		Balance:            1000.00,
		Pin:                "1234",
		TransactionHistory: []Transaction{},
		FailedAttempts:     0,
		IsBlocked:          false,
	}
	card2 := &Card{
		CardNumber:         "6543210987654321",
		Balance:            1000.00,
		Pin:                "1233",
		TransactionHistory: []Transaction{},
		FailedAttempts:     0,
		IsBlocked:          false,
	}

	bank.Accounts[card1.CardNumber] = card1
	bank.Accounts[card2.CardNumber] = card2

	user1 = &User{
		ID:   1,
		Name: "Eshmirza",
		Card: card1,
	}
	user2 := &User{
		ID:   2,
		Name: "Toshmirza",
		Card: card2,
	}
	atm := &ATM{
		Location:      "Morgantown",
		CashAvailable: 5000.00,
		Bank:          bank,
	}
	fmt.Printf("%s ning balansi: %.2f\n", user1.Name, user1.Card.CheckBalance())

}
