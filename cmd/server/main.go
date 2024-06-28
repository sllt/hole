package main

import (
	"flag"
	"github.com/sllt/booby"
	blog "github.com/sllt/booby/log"
	"github.com/sllt/log"
	"hole/internal/server/global"
	"hole/internal/server/router"
	"hole/pkg/logger"
)

func main() {

	addr := flag.String("l", ":3030", "server listen address")
	flag.Parse()
	log.Info("server started...")

	server := booby.NewServer()

	booby.SetLogTag("HOLE")
	blog.SetLogger(logger.New())
	router.RegisterRoutes(server)

	server.Handler.HandleDisconnected(func(client *booby.Client) {
		name, found := client.Get("short-name")
		if found && name.(string) != "" {
			delete(global.Agents, name.(string))
			log.Infof("client %s disconnected", name)
		}

	})

	server.Run(*addr)

}
