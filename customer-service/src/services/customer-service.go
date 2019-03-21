package services

import (
	"customer-service/src/datastore"
	"customer-service/src/messaging"
	customerModel "customer-service/src/models/customers"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// CustomerService - is used to create/udpate/delete customers
// It also publishes events when these operations occur
type CustomerService interface {
	Create(customer customerModel.Customer) (string, error)
}

type customerSvc struct {
	publisher     messaging.Publisher
	customerStore datastore.CustomerStore
}

// NewCustomerService - creates a new instance of CustomerService
func NewCustomerService(publisher messaging.Publisher, customerStore datastore.CustomerStore) CustomerService {
	svc := &customerSvc{
		publisher:     publisher,
		customerStore: customerStore,
	}
	return svc
}

func (svc *customerSvc) Create(customer customerModel.Customer) (string, error) {
	// create customer
	id, err := svc.customerStore.New(customer)
	if err != nil {
		return "", errors.Wrap(err, "Unable to create new customer")
	}
	// publish message
	err = svc.publisher.Publish(newCustomerCreatedEvent(customer))
	if err != nil {
		// We undo by deleting
		delErr := svc.customerStore.Delete(id)
		if delErr != nil {
			logrus.Errorf("Unable to publish customer created. Unable to reverse transaction for id: %s, Error: %s", id, err)
		}
		return "", errors.Wrap(err, "Unable to publish customer created event")
	}
	return id, nil
}
