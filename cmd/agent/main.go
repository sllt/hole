package main

import (
	"github.com/sllt/hole"
	"github.com/sllt/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ServerAddr string
	NpsAddr    string
	Vkey       string
)

func main() {

	log.Info("agent started...")
	hole.StartAgent(ServerAddr)
	hole.StartNpsClient(NpsAddr, Vkey)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	time.Sleep(time.Second)
	log.Info("exiting...")
}
