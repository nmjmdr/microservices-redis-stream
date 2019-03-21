package datastore

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// EventLogStore - data store for event log - last event recorded
type EventLogStore interface {
	Set(logID string) error
	Get() (string, error)
}

type eventLogStore struct {
	db *sql.DB
}

func (i *eventLogStore) Set(logID string) error {
	query := fmt.Sprintf(`UPDATE last_recorded_event_id SET last_event_id = $1`)

	rows, err := i.db.Query(query, logID)
	if err != nil {
		return errors.Wrap(err, "Unable to set last event id in last recorded event id")
	}
	rows.Close()
	return nil
}

func (i *eventLogStore) Get() (string, error) {
	query := fmt.Sprintf(`SELECT last_event_id FROM last_recorded_event_id`)

	rows, err := i.db.Query(query)
	errMessage := "Unable to get last event id in last recorded event id"
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
