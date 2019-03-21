package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"account-service/src/messaging"
)

const listenAddress = ":8090"

func handleExit() {
}

func main() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)

	go func() {
		select {
		case sig := <-signalChan:
			fmt.Printf("Signal received: %v\n", sig)
			// handle graceful close here
			handleExit()
			os.Exit(0)
		}
	}()

	subscriber := messaging.NewRedisStreamSubscriber()
	for {
		subscriber.BlockingListen()
	}

}
