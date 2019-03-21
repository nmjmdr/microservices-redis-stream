package messaging

// GetLastEventLogged - function to get last event id logged
type GetLastEventLogged func() (string, error)

// SetLastEventLogged - function to set last event id logged
type SetLastEventLogged func(logID string) error

// Subscriber - Subscribes to events
type Subscriber interface {
	//BlockingListen(blockTime time.Duration, getFn GetLastEventLogged, setFn SetLastEventLogged) ([]Event, error)
	BlockingListen() ([]Event, error)
}
