package main

import (
	"github.com/sllt/log"
	"hole"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Info("agent started...")
	hole.StartAgent("localhost:3030")

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	time.Sleep(time.Second)
	log.Info("exiting...")
}
