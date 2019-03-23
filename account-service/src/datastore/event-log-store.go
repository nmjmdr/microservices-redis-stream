package datastore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// EventLogStore - data store for event log - last event recorded
type EventLogStore interface {
	Add(logID string, eventType string, payload string) error
	LastEventID() (string, error)
}

type eventLogStore struct {
	db *sql.DB
}

func (i *eventLogStore) Add(logID string, eventType string, payload string) error {
	query := fmt.Sprintf(`INSERT into event_log (event_id, event_type, payload) values ($1, $2, $3)`)

	rows, err := i.db.Query(query, logID, eventType, payload)
	if err != nil {
		return errors.Wrap(err, "Unable to insert into event log")
	}
	rows.Close()
	return nil
}

func (i *eventLogStore) LastEventID() (string, error) {
	query := fmt.Sprintf(`SELECT event_id FROM event_log ORDER BY created_date DESC LIMIT 1`)

	rows, err := i.db.Query(query)
	errMessage := "Unable to get last recorded event id from event log"
	if err != nil {
		return "", errors.Wrap(err, errMessage)
	}
	defer rows.Close()
	logID := ""
	if rows.Next() {
		err := rows.Scan(&logID)
		if err != nil {
			return "", errors.Wrap(err, errMessage)
		}
	}
	return logID, nil
}

// NewEventLogStore - creates a new instance of NewEventLogStore
func NewEventLogStore() (EventLogStore, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to customers database")
	}

	return &eventLogStore{
		db: db,
	}, nil
}
