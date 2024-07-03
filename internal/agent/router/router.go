package router

import (
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/agent/handler"
)

func RegisterAgentRoutes(client booby.Handler) {
	client.Handle("/exit", handler.Exit)
	client.Handle("/exec", handler.ExecuteCmd)
	client.Handle("/socks/start", handler.StartSocksServer)
	client.Handle("/socks/stop", handler.StopSocksServer)
	client.Handle("/shell/start", handler.StartShell)
	client.Handle("/shell/stop", handler.StopShell)
}
