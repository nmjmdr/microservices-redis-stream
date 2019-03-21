package customers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type JsonPurchaseDate struct {
	time.Time
}

// Invoice - represents a invoice object
type Invoice struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	// Ideally should use a composite structure which can representt currency as well
	// we use cents to represent monetary value
	// ideally use https://github.com/shopspring/decimal or big.Float or https://github.com/Rhymond/go-money
	PurchasePriceCents int64            `json:"purchase_price_cents"`
	CreatedDate        time.Time        `json:"created_date"`
	PurchaseDate       JsonPurchaseDate `json:"purchase_date"`
}

// UnmarshalJSON - for Purchase date
func (j *JsonPurchaseDate) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		fmt.Println("Here: ", err, s)
		return err
	}
	j.Time = t
	return nil
}

// MarshalJSON - for Purchase date
func (j JsonPurchaseDate) MarshalJSON() ([]byte, error) {
	val := j.Format("2006-01-02")
	return json.Marshal(val)
}

// Deserialize a Invoice object from the data
func Deserialize(data []byte) (Invoice, error) {
	i := Invoice{}
	err := json.Unmarshal(data, &i)
	if err != nil {
		return Invoice{}, errors.Wrap(err, "Unable to deserialize data as Invoice")
	}
	return i, nil
}

// Serialize a Invoice object from the data
func Serialize(invoice Invoice) ([]byte, error) {
	bytes, err := json.Marshal(invoice)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to serialize invoice")
	}
	return bytes, nil
}

// ValidateInput - validates invoice create input
// ideally should use json schema based validation
func ValidateInput(invoice Invoice) error {
	// Ideally IDs are represented as text. I have used auto-increment ID
	// but ideally should have used GUID
	_, err := strconv.Atoi(invoice.CustomerID)
	if err != nil {
		return errors.Wrap(err, "Unable to convert customer id to int")
	}
	return nil
}
