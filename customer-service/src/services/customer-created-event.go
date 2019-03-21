package services

import customerModel "customer-service/src/models/customers"
import "customer-service/src/messaging"

type customerCreatedEvent struct {
	customer customerModel.Customer
}

func newCustomerCreatedEvent(customer customerModel.Customer) messaging.Event {
	return &customerCreatedEvent{
		customer: customer,
	}
}

func (e *customerCreatedEvent) Type() messaging.EventType {
	return messaging.CustomerCreated
}

func (e *customerCreatedEvent) Payload() string {
	bytes, _ := customerModel.Serialize(e.customer)
	payload := string(bytes)
	return payload
}
