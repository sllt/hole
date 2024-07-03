package handler

import (
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/model/dto"
	"github.com/sllt/hole/internal/server/global"
	"time"
)

func StartAgentShell(ctx *booby.Context) {
	var req dto.ShellRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.Write("param error")
		return
	}

	client := global.Agents[req.AgentName]
	if client == nil {
		ctx.Write("client not found")
		return
	}

	var result string
	client.Conn.Call("/shell/start", nil, &result, time.Second)

	ctx.Write(result)
}

func StopAgentShell(ctx *booby.Context) {
	var req dto.ShellRequest
	if err := ctx.Bind(&req); err != nil {
		ctx.Write("param error")
		return
	}

	client := global.Agents[req.AgentName]
	if client == nil {
		ctx.Write("client not found")
		return
	}

	var result string
	client.Conn.Call("/shell/stop", nil, &result, time.Second)

	ctx.Write(result)
}
