package eventpayloads

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Payload - paylaod when a invoice is created
type Payload struct {
	ID                 string `json:"id"`
	CustomerID         string `json:"customer_id"`
	PurchasePriceCents int64  `json:"purchase_price_cents"`
}

// Deserialize - deserialize payload to InvoiceCreated
func Deserialize(s string) (Payload, error) {
	i := Payload{}
	data := []byte(s)
	err := json.Unmarshal(data, &i)
	if err != nil {
		return Payload{}, errors.Wrap(err, "Unable to deserialize data as InvoiceCreated payload")
	}
	return i, nil
}
