package payment

import "fmt"

// Cash type
type Cash struct {
	Account
}

// CreateCashAccount function
func CreateCashAccount(chargeCh chan float32) *Cash {
	cash := &Cash{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			cash.ProcessPayment(amount)
		}
	}(chargeCh)

	return cash
}

// ProcessPayment method
func (c *Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a cash transaction...")
	return true
}
