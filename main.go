package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	controller := MakeController()
	controller.Start()
	kill := make(chan os.Signal)
	signal.Notify(kill, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-kill
		controller.Stop()
		os.Exit(0)
	}()
	select {}
}
