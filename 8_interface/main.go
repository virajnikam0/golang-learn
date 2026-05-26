package main

import "fmt"

// interface  - what object have and can be do by other after implementation of interface

type RazorPay struct {
}

func (r *RazorPay) razorPayment(amount float32) float32 {
	fmt.Println("payment is done with razor pay", amount)
	return amount + 10.5
}

type Stripe struct {
}

func (s *Stripe) stripePayment(amount float32) float32 {
	fmt.Println("payment is done with stripe pay", amount)
	return amount + 10.5
}

type Payment struct {
	amount float32
}

func (p *Payment) makePayment(amount float32) float32 {
	// pay := RazorPay{}
	pay := Stripe{}
	taxAmount := pay.stripePayment(amount)

	return taxAmount

}

func main() {
	payment_1 := Payment{}
	totalAmount := payment_1.makePayment(100)
	fmt.Println(totalAmount)
}
