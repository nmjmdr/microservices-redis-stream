package services

import (
	"customer-service/src/messaging"
	invoiceModel "customer-service/src/models/invoices"
)

type invoiceCreatedEvent struct {
	invoice invoiceModel.Invoice
}

func newInvoiceCreatedEvent(invoice invoiceModel.Invoice) messaging.Event {
	return &invoiceCreatedEvent{
		invoice: invoice,
	}
}

func (e *invoiceCreatedEvent) Type() messaging.EventType {
	return messaging.InvoiceCreated
}

func (e *invoiceCreatedEvent) Payload() string {
	bytes, _ := invoiceModel.Serialize(e.invoice)
	payload := string(bytes)
	return payload
}
