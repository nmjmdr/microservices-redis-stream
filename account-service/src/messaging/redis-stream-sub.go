package messaging

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

const eventStream = "event-stream"

// EventTypeName - key for stream message which will hold the event type
const EventTypeName = "EventTypeName"

// Payload - key for stream message which will hold the payload
const Payload = "Payload"

type redisStream struct {
	client *redis.Client
}

// NewRedisStreamSubscriber - redis stream as a publisher
func NewRedisStreamSubscriber() Subscriber {
	client := redis.NewClient(&redis.Options{
		Addr:     "svc.redis.io:6379", // ideally fetched from config/env variable
		Password: "",                  // no password set
		DB:       0,                   // use default DB
	})
	return &redisStream{
		client: client,
	}
}

func toStreamMessages(messages []redis.XMessage) []StreamMessage {
	s := []StreamMessage{}
	for _, m := range messages {

		val1, ok1 := m.Values[EventTypeName]
		val2, ok2 := m.Values[Payload]

		if !ok1 || !ok2 {
			continue
		}

		eventType, _ := ParseStringToEventType(val1.(string))
		payload := val2.(string)

		s = append(s, NewStreamMessage(m.ID, NewEvent(eventType, payload)))

	}
	return s
}

func (r *redisStream) BlockingListen(
	count int64,
	blockTime time.Duration,
	lastReadStreamID string,
	readFromStart bool,
) ([]StreamMessage, error) {

	if readFromStart {
		lastReadStreamID = "$"
	}

	streams, err := r.client.XRead(&redis.XReadArgs{
		Streams: []string{eventStream, lastReadStreamID},
		Block:   blockTime,
		Count:   count,
	}).Result()

	if err != nil {
		return []StreamMessage{}, errors.Wrap(err, "Unable not read from event stream")
	}

	if len(streams) == 0 {
		return []StreamMessage{}, errors.New("Could not read from event stream, no results obtained")
	}

	if len(streams[0].Messages) == 0 {
		return []StreamMessage{}, nil
	}

	return toStreamMessages(streams[0].Messages), nil
}
