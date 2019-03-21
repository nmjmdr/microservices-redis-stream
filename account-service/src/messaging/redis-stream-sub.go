package messaging

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

const eventStream = "event-stream"

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

// func (r *redisStream) BlockingListen(blockTime time.Duration, getFn GetLastEventLogged, setFn SetLastEventLogged) ([]Event, error) {
func (r *redisStream) BlockingListen() ([]Event, error) {
	result, err := r.client.XRead(&redis.XReadArgs{
		Streams: []string{eventStream, "$"},
	}).Result()
	fmt.Println("-----> ", result, err)
	/*
		[{event-stream [{1553156572229-0 map[InvoiceCreated:{"id":"","customer_id":"1","purchase_price_cents":15000,"created_date":"0001-01-01T00:00:00Z","purchase_date":"2019-01-03"}]}]}]
	*/
	if err != nil {
		return []Event{}, errors.Wrap(err, "Unable to publish event to event stream")
	}
	return []Event{}, nil
}
