package eventsink

import (
	"account-service/src/datastore"
	"account-service/src/messaging"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// TO DO: read from config
const blockTime = 1000 * time.Millisecond

// SyncMessageFn - function to sync message to db and process it
type SyncMessageFn func(message messaging.StreamMessage) error

// Process - processes stream messages in a continous loop
func Process(quitChan chan bool,
	subscriber messaging.Subscriber,
	logStore datastore.EventLogStore,
	syncFn SyncMessageFn,
) {
	lastReadID, err := logStore.LastEventID()
	if err != nil {
		logrus.Errorf("Unable to read last read id, Error: %s", err)
		return
	}

loop:
	for {
		select {
		case <-quitChan:
			break loop
		default:
			// Process messages one by one
			messages, err := subscriber.BlockingListen(1, blockTime, lastReadID, (len(lastReadID) == 0))
			if err != nil {
				logrus.Errorf("Unable to listen to stream events, Error: %s", err)
				return
			}
			if len(messages) == 0 {
				logrus.Info("Nothing to read from stream")
				continue
			}
			fmt.Println("Event got: ", messages[0].ID())
			err = syncFn(messages[0])
			if err != nil {
				logrus.Errorf("Unable to sync stream events to database, Error: %s", err)
				return
			}
			err = logStore.Add(messages[0].ID(), messages[0].Type().String(), messages[0].Payload())
			if err != nil {
				logrus.Errorf("Unable to set last read stream event ID, Error: %s", err)
				return
			}
			lastReadID = messages[0].ID()
		}
	}
}
