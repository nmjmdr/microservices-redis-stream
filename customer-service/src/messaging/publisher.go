package messaging

// Publisher - Publishes events
type Publisher interface {
	Publish(evt Event) error
}
