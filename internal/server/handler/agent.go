package handler

import (
	"github.com/sllt/booby"
	"hole/internal/model"
	"hole/internal/server/global"
)

func AgentList(ctx *booby.Context) {
	resp := make([]*model.Agent, 0)

	for _, v := range global.Agents {
		resp = append(resp, v)
	}

	_ = ctx.Write(resp)
}

func AgentRegister(ctx *booby.Context) {
	var agent model.Agent

	err := ctx.Bind(&agent)
	if err != nil {
		return
	}
	agent.Conn = ctx.Client

	global.Agents[agent.ShortName] = &agent
	ctx.Client.Set("short-name", agent.ShortName)
	_ = ctx.Write("ok")
}
