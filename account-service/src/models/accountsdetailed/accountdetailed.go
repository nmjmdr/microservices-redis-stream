package accountsdetailed

import (
	"account-service/src/models/accounts"
	"account-service/src/models/linkedcustomers"
	"encoding/json"

	"github.com/pkg/errors"
)

// AccountDetailed - represents an Account object along with
type AccountDetailed struct {
	Account           accounts.Account                 `json:"account"`
	Customers         []linkedcustomers.LinkedCustomer `json:"customers"`
	TotalRevenueCents int64                            `json:"total_revenue_cents"`
}

func sumRevenues(customers []linkedcustomers.LinkedCustomer) int64 {
	total := int64(0)
	for _, c := range customers {
		total = total + c.TotalRevenueCents
	}
	return total
}

// NewAccountDetailed - creates a new instance of `account detailed`
func NewAccountDetailed(account accounts.Account,
	customers []linkedcustomers.LinkedCustomer) AccountDetailed {

	totalRevenue := sumRevenues(customers)
	return AccountDetailed{
		Account:           account,
		Customers:         customers,
		TotalRevenueCents: totalRevenue,
	}
}

// Serialize an Account object from the data
func Serialize(accountDetailed AccountDetailed) ([]byte, error) {
	bytes, err := json.Marshal(accountDetailed)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to serialize Account")
	}
	return bytes, nil
}
