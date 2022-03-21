package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Press CTRL + C to interrupt")

	select {
	case <-signalChannel:
		fmt.Println("Signal handled")
		os.Exit(0)
	}
}
