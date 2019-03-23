package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"account-service/src/synchronizer"

	"github.com/sirupsen/logrus"
)

const listenAddress = ":8090"

var sync synchronizer.Sync

func handleExit() {
	sync.Stop()
}

func main() {
	var err error
	sync, err = synchronizer.NewSync()
	if err != nil {
		logrus.Errorf("Unable to start event synchronizer, Error: %s", err)
		os.Exit(1)
	}
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)

	fmt.Println("Calling sync")
	sync.Start()
	fmt.Println("Called sync")
	select {
	case sig := <-signalChan:
		fmt.Printf("Signal received: %v\n", sig)
		handleExit()
		os.Exit(0)
	}

}
