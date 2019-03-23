package datastore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// LinkedCustomerStore - data store for linked customers
type LinkedCustomerStore interface {
	New(customerID string, accountID int, name string) error
	AddRevenue(customerID string, deltaRevenueCents int64) (int64, error)
	SubtractRevenue(customerID string, deltaRevenueCents int64) (int64, error)
	Delete(customerID string) error
}

type linkedCustomerStore struct {
	db *sql.DB
}

func (i *linkedCustomerStore) New(customerID string, accountID int, name string) error {
	query := fmt.Sprintf(`INSERT INTO linked_customers (customer_id, account_id) values ($1, $2, $3)`)

	rows, err := i.db.Query(query, customerID, accountID, name)
	if err != nil {
		return errors.Wrap(err, "Unable to save linked customer to database")
	}
	rows.Close()
	return nil
}

type operation int

const (
	add operation = iota
	sub
)

func (o operation) String() string {
	return [...]string{"add", "subtract"}[o]
}

func (i *linkedCustomerStore) deltaRevenue(customerID string, delta int64, op operation) (int64, error) {
	query := fmt.Sprintf(`UPDATE linked_customers SET total_revenue = total_revenue + $2 WHERE customer_id = $1 RETURNING total_revenue`)
	if op == sub {
		query = fmt.Sprintf(`UPDATE linked_customers SET total_revenue = total_revenue - $2 WHERE customer_id = $1 RETURNING total_revenue`)
	}

	rows, err := i.db.Query(query, customerID, delta)
	if err != nil {
		return 0, errors.Wrapf(err, "Unable to %s delta revenue to linked customer", op.String())
	}
	defer rows.Close()
	totalRev := int64(0)
	if rows.Next() {
		err := rows.Scan(&totalRev)
		if err != nil {
			return 0, errors.Wrap(err, "Unable to read total revenue after delta operation")
		}
	} else {
		return 0, errors.Wrap(err, "Unable to read total revenue after delta operation")
	}
	return totalRev, nil
}

func (i *linkedCustomerStore) AddRevenue(customerID string, delta int64) (int64, error) {
	return i.deltaRevenue(customerID, delta, add)
}

func (i *linkedCustomerStore) SubtractRevenue(customerID string, delta int64) (int64, error) {
	return i.deltaRevenue(customerID, delta, sub)
}

// Delete - delete an invoice
func (i *linkedCustomerStore) Delete(id string) error {
	return nil
}

// NewLinkedCustomerStore - creates a new instance of LinkedCustomerStore store
func NewLinkedCustomerStore() (LinkedCustomerStore, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to customers database")
	}

	return &linkedCustomerStore{
		db: db,
	}, nil
}
