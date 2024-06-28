package router

import (
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/client/handler"
)

func RegisterClientRoutes(server *booby.Server) {

	server.Handler.Handle("/exit", handler.Exit)
	server.Handler.Handle("/exec", handler.Execute)
	server.Handler.Handle("/db", handler.ExecDB)
}
