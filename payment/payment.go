package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// CreditCard type
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// CreateCreditAccount function
func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

// ProcessPayment method
func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a credit card payment...")
	return true
}

// OwnerName method
func (c CreditCard) OwnerName() string {
	return c.ownerName
}

// SetOwnerName method
func (c *CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid owner name provided")
	}
	c.ownerName = value
	return nil
}

// CardNumber method
func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

// SetCardNumber method
func (c *CreditCard) SetCardNumber(value string) error {
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid credit card number format")
	}
	c.cardNumber = value
	return nil
}

// ExpirationDate method
func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

// SetExpirationDate method
func (c *CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date must lie in the future")
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

// SecurityCode method
func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

// SetSecurityCode method
func (c *CreditCard) SetSecurityCode(value int) {
	if value < 100 || value > 999 {

	}
}

// AvailableCredit method
func (c CreditCard) AvailableCredit() float32 {
	return 5000.0
}
