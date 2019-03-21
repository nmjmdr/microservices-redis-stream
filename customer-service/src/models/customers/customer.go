package customers

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Customer - represents a customer object
type Customer struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"created_date"`
}

// Deserialize a customer object from the data
func Deserialize(data []byte) (Customer, error) {
	c := Customer{}
	err := json.Unmarshal(data, &c)
	if err != nil {
		return Customer{}, errors.Wrap(err, "Unable to deserialize data as customer")
	}
	return c, nil
}

// Serialize a customer object from the data
func Serialize(customer Customer) ([]byte, error) {
	bytes, err := json.Marshal(customer)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to serialize customer")
	}
	return bytes, nil
}

// ValidateInput - validates customer create input
// ideally should use json schema based validation
func ValidateInput(customer Customer) error {
	if len(customer.Name) == 0 || len(customer.AccountID) == 0 {
		return errors.New("Properties `Name` and `AccountID` should be set for a customer")
	}
	return nil
}
