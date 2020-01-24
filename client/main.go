package main

import (
	"fmt"

	"github.com/robertwitt/go-tutorial-gopherpay/payment"
)

func main() {
	var option payment.PaymentOption

	option = payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2021,
		123,
	)
	option.ProcessPayment(500)

	chargeCh := make(chan float32)
	cash := payment.CreateCashAccount(chargeCh)
	chargeCh <- 500
	var a string
	fmt.Scanln(&a)

	cash.AvailableFunds()
}
