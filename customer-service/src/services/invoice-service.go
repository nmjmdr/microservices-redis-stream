package services

import (
	"customer-service/src/datastore"
	"customer-service/src/messaging"
	invoiceModel "customer-service/src/models/invoices"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// InvoiceService - is used to create/udpate/delete invoices
// It also publishes events when these operations occur
type InvoiceService interface {
	Create(invoice invoiceModel.Invoice) (string, error)
}

type invoiceSvc struct {
	publisher    messaging.Publisher
	invoiceStore datastore.InvoiceStore
}

// NewInvoiceService - creates a new instance of NewInvoicerService
func NewInvoiceService(publisher messaging.Publisher, invoiceStore datastore.InvoiceStore) InvoiceService {
	svc := &invoiceSvc{
		publisher:    publisher,
		invoiceStore: invoiceStore,
	}
	return svc
}

func (svc *invoiceSvc) Create(invoice invoiceModel.Invoice) (string, error) {
	// create invoice
	id, err := svc.invoiceStore.New(invoice)
	if err != nil {
		return "", errors.Wrap(err, "Unable to create new invoice")
	}
	// publish message
	invoice.ID = id
	err = svc.publisher.Publish(newInvoiceCreatedEvent(invoice))
	if err != nil {
		// We undo by deleting
		delErr := svc.invoiceStore.Delete(id)
		if delErr != nil {
			logrus.Errorf("Unable to publish invoice created. Unable to reverse transaction for id: %s, Error: %s", id, err)
		}
		return "", errors.Wrap(err, "Unable to publish invoice created event")
	}
	return id, nil
}
