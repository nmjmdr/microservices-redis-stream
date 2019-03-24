package datastore

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// AuthStore - data store for user details
type AuthStore interface {
	Match(username string, hashedPass string) (bool, error)
}

type authStore struct {
	db *sql.DB
}

func (c *accountStore) Match(username string, hashedPass string) (bool, error) {
	query := `SELECT EXISTS (SELECT username FROM users WHERE username = $1 and password_hash = $2);`
	rows, err := c.db.Query(query, username, hashedPass)
	if err != nil {
		return false, errors.Wrap(err, "Unable to get login details")
	}
	defer rows.Close()
	result := false
	if rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			return false, errors.Wrap(err, "Unable to read result while trying to get login details")
		}
	} else {
		return false, errors.Wrap(err, "No rows returned while trying to get login details")
	}
	return result, nil
}

// NewAuthStore - creates a new instance of AuthStore
func NewAuthStore() (AuthStore, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to database")
	}

	return &accountStore{
		db: db,
	}, nil
}
