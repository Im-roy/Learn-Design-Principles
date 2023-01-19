package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float32) string
}

type CreditCard struct {
	Name     string
	CardNo   string
	CVV      string
	ExpMonth string
	ExpYear  string
}

func (c *CreditCard) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using credit/debit card", amount)
}

type PayPal struct {
	Email    string
	Password string
}

func (p *PayPal) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using PayPal", amount)
}

type Item struct {
	Name     string
	Quantity int
	Price    float32
	Payment  PaymentStrategy
}

func (i *Item) Total() float32 {
	return float32(i.Quantity) * i.Price
}

func (i *Item) Pay() string {
	return i.Payment.Pay(i.Total())
}

func main() {
	item := &Item{
		Name:     "book",
		Quantity: 1,
		Price:    25.0,
	}
	item.Payment = &PayPal{
		Email:    "test@example.com",
		Password: "testpassword",
	}
	fmt.Println(item.Pay())
	item.Payment = &CreditCard{
		Name:     "John Doe",
		CardNo:   "1234123412341234",
		CVV:      "123",
		ExpMonth: "01",
		ExpYear:  "25",
	}
	fmt.Println(item.Pay())
}
