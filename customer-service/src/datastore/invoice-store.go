package datastore

import (
	invoiceModel "customer-service/src/models/invoices"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// InvoiceStore - data store for invoices
type InvoiceStore interface {
	New(invoice invoiceModel.Invoice) (string, error)
	Delete(id string) error
}

type invoiceStore struct {
	db *sql.DB
}

func (i *invoiceStore) New(invoice invoiceModel.Invoice) (string, error) {
	query := fmt.Sprintf(`INSERT INTO invoice (customer_id, purchase_date, purchase_price) values ($1, $2, $3) RETURNING id`)
	customerID, err := strconv.Atoi(invoice.CustomerID)
	if err != nil {
		return "", errors.Wrap(err, "Unable to convert customer id to int")
	}

	purchaseDate := invoice.PurchaseDate.Format("2006/01/02")
	rows, err := i.db.Query(query, customerID, purchaseDate, invoice.PurchasePriceCents)
	if err != nil {
		return "", errors.Wrap(err, "Unable to save invoice to database")
	}
	defer rows.Close()
	id := ""
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return "", errors.Wrap(err, "Unable to read invoice id from the saved record")
		}
	} else {
		return "", errors.Wrap(err, "Unable to read invoice id from the saved record")
	}
	return id, nil
}

// Delete - delete an invoice
func (i *invoiceStore) Delete(id string) error {
	return nil
}

// NewInvoiceStore - creates a new instance of invoice store
func NewInvoiceStore() (InvoiceStore, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to customers database")
	}

	return &invoiceStore{
		db: db,
	}, nil
}
