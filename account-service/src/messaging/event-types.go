package messaging

// Ideally messaging is a common library, TO DO: move to common library
// The event names are shared between services

// EventType - type of the event being published
type EventType int

const (
	// CustomerCreated - published when a new customer is created
	CustomerCreated EventType = iota
	// InvoiceCreated - published when a new invoice is created
	InvoiceCreated
)

func (e EventType) String() string {
	return [...]string{"CustomerCreated", "InvoiceCreated"}[e]
}
