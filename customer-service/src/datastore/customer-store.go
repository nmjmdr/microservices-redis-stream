package datastore

import (
	customerModel "customer-service/src/models/customers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// CustomerStore - data store for customers
type CustomerStore interface {
	New(customer customerModel.Customer) (string, error)
	Delete(id string) error
}

type customerStore struct {
	db *sql.DB
}

func (c *customerStore) New(customer customerModel.Customer) (string, error) {
	query := fmt.Sprintf(`INSERT INTO customer (name, account_id) values ($1, $2) RETURNING id`)
	rows, err := c.db.Query(query, customer.Name, customer.AccountID)
	if err != nil {
		return "", errors.Wrap(err, "Unable to save customer to database")
	}
	defer rows.Close()
	id := ""
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return "", errors.Wrap(err, "Unable to read customer id from the saved record")
		}
	} else {
		return "", errors.Wrap(err, "Unable to read customer id from the saved record")
	}
	return id, nil
}

// Delete - delete a customer
func (c *customerStore) Delete(id string) error {
	return nil
}

// NewCustomerStore - creates a new instance of customers store
func NewCustomerStore() (CustomerStore, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to customers database")
	}

	return &customerStore{
		db: db,
	}, nil
}
