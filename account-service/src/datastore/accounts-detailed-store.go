package datastore

import (
	accountsDetailed "account-service/src/models/accountsdetailed"
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

func (c *accountStore) GetAll() ([]accountsDetailed.AccountDetailed, error) {
	query := fmt.Sprintf(`SELECT a.id account_id, 
	a.name account_name, 
	a.description account_description, 
	a.owned_by account_owned_by, 
	c.customer_id customer_id, 
	c.name customer_name, 
	c.total_revenue customer_revenue 
	FROM accounts a 
	INNER JOIN 
	linked_customers c 
	ON a.id = c.account_id 
	ORDER BY a.id;`)

	rows, err := c.db.Query(query)
	if err != nil {
		return []accountsDetailed.AccountDetailed{}, errors.Wrap(err, "Unable to get all account details")
	}
	results := []accountsDetailed.AccountDetailed{}
	defer rows.Close()

	return results, nil
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
