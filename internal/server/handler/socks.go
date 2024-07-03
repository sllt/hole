package handler

import (
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/model/dto"
	"github.com/sllt/hole/internal/server/global"
	"time"
)

func StartAgentSocks(ctx *booby.Context) {
	var req dto.StartSocksRequest
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
	client.Conn.Call("/socks/start", &req, &result, time.Second)

	ctx.Write(result)
}

func StopAgentSocks(ctx *booby.Context) {
	var req dto.StartSocksRequest
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
	client.Conn.Call("/socks/stop", &req, &result, time.Second)

	ctx.Write(result)
}
