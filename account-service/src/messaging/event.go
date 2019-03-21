package messaging

// Event - a business event
type Event interface {
	Type() EventType
	Payload() string
}

type event struct {
	eventType EventType
	payload   string
}

func (e *event) Type() EventType {
	return e.eventType
}

func (e *event) Payload() string {
	return e.payload
}

// NewEvent - creates a new event
func NewEvent(eventType EventType, payload string) Event {
	return &event{
		eventType: eventType,
		payload:   payload,
	}
}
