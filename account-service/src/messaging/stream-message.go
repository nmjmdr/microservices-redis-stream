package messaging

// StreamMessage - a message obtained from the stream
type StreamMessage interface {
	ID() string
	Event
}

type streamMessage struct {
	id  string
	evt Event
}

func (s *streamMessage) ID() string {
	return s.id
}

func (s *streamMessage) Type() EventType {
	return s.evt.Type()
}

func (s *streamMessage) Payload() string {
	return s.evt.Payload()
}

// NewStreamMessage - returns a new stream message
func NewStreamMessage(id string, evt Event) StreamMessage {
	return &streamMessage{
		id:  id,
		evt: evt,
	}
}
