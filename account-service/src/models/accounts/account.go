package accounts

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Account - represents an Account object
type Account struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnedBy     string    `json:"owned_by"`
	CreatedDate time.Time `json:"created_date"`
}

// Deserialize an Account object from the data
func Deserialize(data []byte) (Account, error) {
	c := Account{}
	err := json.Unmarshal(data, &c)
	if err != nil {
		return Account{}, errors.Wrap(err, "Unable to deserialize data as an Account")
	}
	return c, nil
}

// Serialize an Account object from the data
func Serialize(account Account) ([]byte, error) {
	bytes, err := json.Marshal(account)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to serialize Account")
	}
	return bytes, nil
}

// ValidateInput - validates an Account create input
// ideally should use json schema based validation
func ValidateInput(account Account) error {
	if len(account.Name) == 0 || len(account.OwnedBy) == 0 {
		return errors.New("Properties `Name` and `OwnedBy` should be set for an Account")
	}
	return nil
}
