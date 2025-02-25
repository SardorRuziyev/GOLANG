/*
	E-commerce to‘lov tizimi

Vazifa:
To‘lov tizimini boshqarish uchun PaymentProcessor interfeysini yarating, unda quyidagi metodlar bo‘lsin:

Pay(amount float64) string
Refund(amount float64) string
TransactionHistory() []string
Keyin, ushbu interfeysni implement qiluvchi quyidagi to‘lov tizimlarini yozing:

CreditCardProcessor: kartadan to‘lovlar qabul qiladi.
PayPalProcessor: PayPal orqali to‘lovlarni boshqaradi.
CryptoProcessor: kriptovalyutalar orqali to‘lovlarni amalga oshiradi.
Har bir tizim o‘ziga xos xabar va tranzaksiya tarixini qaytarsin. Foydalanuvchi tanlagan to‘lov turini boshqaradigan dastur yozing.
*/
package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount float64) string
	Refund(amount float64) string
	TransactionHistory() []string
}

type CreditCardProcessor struct {
	transactions []string
}
type PayPalProcessor struct {
	transactions []string
}
type CryptoProcessor struct {
	transactions []string
}

func (c *CreditCardProcessor) Pay(amount float64) string {
	transaction := fmt.Sprintf("CreditCardProcessor: %.2f summa to'lov qilindi.", amount)
	c.transactions = append(c.transactions, transaction)
	return transaction
}
func (c *CreditCardProcessor) Refund(amount float64) string {
	transaction := fmt.Sprintf("CreditCardProcessor: %.2f summa qaytarildi", amount)
	c.transactions = append(c.transactions, transaction)
	return transaction
}
func (c *CreditCardProcessor) TransactionHistory() []string {
	return c.transactions
}

func (p *PayPalProcessor) Pay(amount float64) string {
	transaction := fmt.Sprintf("Paypal: %.2f summa to'lov qilindi.", amount)
	p.transactions = append(p.transactions, transaction)
	return transaction
}
func (p *PayPalProcessor) Refund(amount float64) string {
	transaction := fmt.Sprintf("Paypal: %.2f summa qaytarildi", amount)
	p.transactions = append(p.transactions, transaction)
	return transaction
}
func (p *PayPalProcessor) TransactionHistory() []string {
	return p.transactions
}

func (cp *CryptoProcessor) Pay(amount float64) string {
	transaction := fmt.Sprintf("Crypto Wallet: %.2f summa to'lov qilindi.", amount)
	cp.transactions = append(cp.transactions, transaction)
	return transaction
}
func (cp *CryptoProcessor) Refund(amount float64) string {
	transaction := fmt.Sprintf("Crypto Wallet: %.2f summa qaytarildi", amount)
	cp.transactions = append(cp.transactions, transaction)
	return transaction
}
func (cp *CryptoProcessor) TransactionHistory() []string {
	return cp.transactions
}

func main() {
	var processor PaymentProcessor
	fmt.Println("To'lov turini tanlang")
	fmt.Println("1. Credit card")
	fmt.Println("2. Paypal")
	fmt.Println("3. Crypto Wallet")
	fmt.Println("Exit")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		processor = &CreditCardProcessor{}
	case 2:
		processor = &PayPalProcessor{}
	case 3:
		processor = &CryptoProcessor{}
	default:
		fmt.Println("Noto'g'ri tanlov!")
		return
	}
	fmt.Println(processor.Pay(100.50))

	fmt.Println(processor.Refund(50.25))

	fmt.Println("Tranzaksiya tarixi:")
	for _, transaction := range processor.TransactionHistory() {
		fmt.Println(transaction)
	}

}
