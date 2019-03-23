package customercreated

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Payload - paylaod when a customer is created
type Payload struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
}

// Deserialize - deserialize payload to InvoiceCreated
func Deserialize(s string) (Payload, error) {
	i := Payload{}
	data := []byte(s)
	err := json.Unmarshal(data, &i)
	if err != nil {
		return Payload{}, errors.Wrap(err, "Unable to deserialize data as CustomerCreated payload")
	}
	return i, nil
}
