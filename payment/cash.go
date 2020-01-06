package payment

import "fmt"

// Cash type
type Cash struct{}

// CreateCashAccount function
func CreateCashAccount() *Cash {
	return &Cash{}
}

// ProcessPayment method
func (c *Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a cash transaction...")
	return true
}
