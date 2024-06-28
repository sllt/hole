package router

import (
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/server/handler"
)

func RegisterRoutes(server *booby.Server) {

	server.Handler.Handle("/ping", handler.Ping)
	server.Handler.Handle("/agent/list", handler.AgentList)
	server.Handler.Handle("/agent/register", handler.AgentRegister)

	// handle
	server.Handler.Handle("/handle/exit", handler.HandleExit)
	server.Handler.Handle("/handle/exec", handler.HandleExecute)
}
