package messaging

import "time"

// Subscriber - Subscribes to events
type Subscriber interface {
	BlockingListen(
		count int64,
		blockTime time.Duration,
		lastReadStreamID string,
		readFromStart bool,
	) ([]StreamMessage, error)
}
