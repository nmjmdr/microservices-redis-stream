package synchronizer

import (
	"account-service/src/datastore"
	"account-service/src/eventsink"
	"account-service/src/messaging"
	customercreated "account-service/src/models/eventpayloads/customercreated"
	invoicecreated "account-service/src/models/eventpayloads/invoicecreated"
	"strconv"

	"github.com/pkg/errors"
)

type Sync interface {
	Start() error
	Stop()
}
type sync struct {
	stop                chan bool
	subscriber          messaging.Subscriber
	logStore            datastore.EventLogStore
	linkedCustomerStore datastore.LinkedCustomerStore
}

// NewSync - provides a new sync implementation
func NewSync() (Sync, error) {
	logStore, err := datastore.NewEventLogStore()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to Sync events, cannot connect to log store")
	}
	linkedCustomerStore, err := datastore.NewLinkedCustomerStore()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to Sync events, cannot connect to linked customer store")
	}
	return &sync{
		stop:                make(chan bool),
		subscriber:          messaging.NewRedisStreamSubscriber(),
		logStore:            logStore,
		linkedCustomerStore: linkedCustomerStore,
	}, nil
}

func (s *sync) syncMessage(message messaging.StreamMessage) error {
	// process message here
	switch message.Type() {
	case messaging.CustomerCreated:
		createdCustomer, err := customercreated.Deserialize(message.Payload())
		if err != nil {
			return errors.Wrap(err, "Unable to deserialize customer created event")
		}
		accountID, err := strconv.Atoi(createdCustomer.AccountID)
		if err != nil {
			return errors.Wrap(err, "Unable to sync customer created event to database, payload account ID is not an integer")
		}
		err = s.linkedCustomerStore.New(createdCustomer.ID, accountID)
		if err != nil {
			return errors.Wrap(err, "Unable to save customer created event to database")
		}
	case messaging.InvoiceCreated:
		createdInvoice, err := invoicecreated.Deserialize(message.Payload())
		if err != nil {
			return errors.Wrap(err, "Unable to deserialize invoice created event")
		}
		_, err = s.linkedCustomerStore.AddRevenue(createdInvoice.CustomerID, createdInvoice.PurchasePriceCents)
		if err != nil {
			return errors.Wrap(err, "Unable to add revene when an invoice was created")
		}
	}
	return nil
}

func (s *sync) Start() error {
	go func() {
		eventsink.Process(s.stop, s.subscriber, s.logStore, s.syncMessage)
	}()
	return nil
}

func (s *sync) Stop() {
	s.stop <- true
}
