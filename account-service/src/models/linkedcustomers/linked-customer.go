package linkedcustomers

import (
	"strconv"

	"github.com/pkg/errors"
)

// LinkedCustomer - represents a customer linked to an account, also stores the relevant statistics such as total revenue for a customer
type LinkedCustomer struct {
	CustomerID string `json:"customer_id"`
	// Ideally should use a composite structure which can representt currency as well
	// we use cents to represent monetary value
	// ideally use https://github.com/shopspring/decimal or big.Float or https://github.com/Rhymond/go-money
	TotalRevenueCents int64  `json:"total_revenue_cents"`
	Name              string `json:"name"`
}

// ValidateInput - validates invoice create input
// ideally should use json schema based validation
func ValidateInput(customerID string) error {
	// Ideally IDs are represented as text. I have used auto-increment ID
	// but ideally should have used GUID
	_, err := strconv.Atoi(customerID)
	if err != nil {
		return errors.Wrap(err, "Unable to convert customer id to int")
	}
	return nil
}
