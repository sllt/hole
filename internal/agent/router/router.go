package router

import (
	"github.com/sllt/booby"
	"hole/internal/agent/handler"
)

func RegisterAgentRoutes(client booby.Handler) {
	client.Handle("/exit", handler.Exit)
	client.Handle("/exec", handler.ExecuteCmd)
}
