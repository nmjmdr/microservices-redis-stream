package messaging

// Event - a business event
type Event interface {
	Type() EventType
	Payload() string
}
