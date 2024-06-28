package handler

import (
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/model/dto"
	"github.com/sllt/hole/internal/server/global"
	"time"
)

func HandleExit(ctx *booby.Context) {
	var req dto.ExitRequest

	err := ctx.Bind(&req)
	if err != nil {
		ctx.Write(err)
		return
	}

	client := global.Agents[req.AgentName]
	if client == nil {
		ctx.Write("client not found")
		return
	}

	client.Conn.Call("/exit", nil, nil, time.Second)
	_ = ctx.Write("ok")
}

func HandleExecute(ctx *booby.Context) {
	var req dto.ExecuteRequest

	err := ctx.Bind(&req)
	if err != nil {
		ctx.Write(err)
		return
	}

	client := global.Agents[req.AgentName]
	if client == nil {
		ctx.Write("client not found")
		return
	}

	var result string
	err = client.Conn.Call("/exec", req, &result, time.Second)

	if err != nil {
		_ = ctx.Write(err.Error())
		return
	}
	_ = ctx.Write(result)
}
