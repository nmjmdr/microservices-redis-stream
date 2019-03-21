package datastore

import (
	accountModel "account-service/src/models/accounts"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// AccountStore - data store for accounts
type AccountStore interface {
	New(account accountModel.Account) (string, error)
	Delete(id string) error
}

type accountStore struct {
	db *sql.DB
}

func (c *accountStore) New(account accountModel.Account) (string, error) {
	query := fmt.Sprintf(`INSERT INTO accounts (name, description, owned_by) values ($1, $2, $3) RETURNING id`)
	rows, err := c.db.Query(query, account.Name, account.Description, account.OwnedBy)
	if err != nil {
		return "", errors.Wrap(err, "Unable to save account to database")
	}
	defer rows.Close()
	id := ""
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return "", errors.Wrap(err, "Unable to read account id from the saved record")
		}
	} else {
		return "", errors.Wrap(err, "Unable to read account id from the saved record")
	}
	return id, nil
}

// Delete - delete an account
func (c *accountStore) Delete(id string) error {
	return nil
}

// NewAccountStore - creates a new instance of accouunt store
func NewAccountStore() (AccountStore, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to accounts database")
	}

	return &accountStore{
		db: db,
	}, nil
}
