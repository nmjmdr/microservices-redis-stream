package main

import (
	"customer-service/src/router"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	
)

const listenAddress = ":8080"

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

	fmt.Printf("API listening on: %s ...", listenAddress)
	fmt.Println()

	router.Start(listenAddress)

}
