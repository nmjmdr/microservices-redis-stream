package datastore

import (
	accountsModel "account-service/src/models/accounts"
	accountsDetailed "account-service/src/models/accountsdetailed"
	linkedcustomers "account-service/src/models/linkedcustomers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// AccountsDetailedStore - data store for `account detailed`
type AccountsDetailedStore interface {
	GetAll() ([]accountsDetailed.AccountDetailed, error)
}

type accountsDetailedStore struct {
	db *sql.DB
}

func toArray(m map[int]*accountsDetailed.AccountDetailed) []accountsDetailed.AccountDetailed {
	results := []accountsDetailed.AccountDetailed{}
	for _, v := range m {
		results = append(results, *v)
	}
	return results
}

func (c *accountStore) GetAll() ([]accountsDetailed.AccountDetailed, error) {
	query := fmt.Sprintf(`SELECT a.id account_id, 
	a.name account_name, 
	a.description account_description, 
	a.owned_by account_owned_by, 
	COALESCE(c.customer_id, '') customer_id, 
	COALESCE(c.name,'') customer_name, 
	COALESCE(c.total_revenue,0) customer_revenue 
	FROM accounts a 
	LEFT JOIN 
	linked_customers c 
	ON a.id = c.account_id 
	ORDER BY a.id;`)

	rows, err := c.db.Query(query)
	if err != nil {
		return []accountsDetailed.AccountDetailed{}, errors.Wrap(err, "Unable to get all account details")
	}
	results := make(map[int]*accountsDetailed.AccountDetailed)
	defer rows.Close()

	var accountID int
	var accountName, customerName, ownedBy, accountDescription, customerID string
	var customerRevenue int64

	for rows.Next() {
		err := rows.Scan(&accountID, &accountName, &accountDescription, &ownedBy, &customerID, &customerName, &customerRevenue)
		if err != nil {
			return []accountsDetailed.AccountDetailed{}, errors.Wrap(err, "Unable to get all account details")
		}
		_, ok := results[accountID]
		if !ok {
			results[accountID] = &accountsDetailed.AccountDetailed{
				Account: accountsModel.Account{
					ID:          accountID,
					Name:        accountName,
					Description: accountDescription,
					OwnedBy:     ownedBy,
				},
				Customers:         []linkedcustomers.LinkedCustomer{},
				TotalRevenueCents: 0,
			}
		}
		if len(customerID) > 0 {
			customer := linkedcustomers.LinkedCustomer{
				CustomerID:        customerID,
				Name:              customerName,
				TotalRevenueCents: customerRevenue,
			}
			results[accountID].Customers = append(results[accountID].Customers, customer)
			results[accountID].TotalRevenueCents += customer.TotalRevenueCents
		}
	}

	return toArray(results), nil
}

// NewAccountsDetailedStore - creates a new instance of accouunt detailed store
func NewAccountsDetailedStore() (AccountsDetailedStore, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to accounts database")
	}
	return &accountStore{
		db: db,
	}, nil
}
