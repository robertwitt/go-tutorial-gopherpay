package payment

import (
	"fmt"
)

// Account type
type Account struct{}

// AvailableFunds method
func (a *Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds")
	return 0
}

// ProcessPayment method
func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}
