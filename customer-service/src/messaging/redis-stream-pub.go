package messaging

import (
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

// NewRedisStreamPublisher - redis stream as a publisher
func NewRedisStreamPublisher() Publisher {
	client := redis.NewClient(&redis.Options{
		Addr:     "svc.redis.io:6379", // ideally fetched from config/env variable
		Password: "",                  // no password set
		DB:       0,                   // use default DB
	})
	return &redisStream{
		client: client,
	}
}

func (r *redisStream) Publish(evt Event) error {
	vals := make(map[string]interface{})
	vals[EventTypeName] = evt.Type().String()
	vals[Payload] = evt.Payload()
	_, err := r.client.XAdd(&redis.XAddArgs{
		Stream: eventStream,
		Values: vals,
	}).Result()
	if err != nil {
		return errors.Wrapf(err, "Unable to publish event %s to event stream", evt.Type().String())
	}
	return nil
}
