package service

import (
	"github.com/sllt/hole/internal/client/global"
	"github.com/sllt/hole/internal/model"
	"github.com/sllt/hole/internal/model/dto"
	"time"
)

type client struct {
}

var Client = &client{}

func (c *client) GetAgentList() ([]*model.Agent, error) {
	agents := make([]*model.Agent, 0)
	err := global.Client.Call("/agent/list", nil, &agents, time.Second)
	if err != nil {
		return nil, err
	}

	return agents, nil
}

func (c *client) Exit(agentName string) (string, error) {
	req := &dto.ExitRequest{
		AgentName: agentName,
	}

	var result string

	err := global.Client.Call("/handle/exit", req, &result, time.Second)
	return result, err
}

func (c *client) Exec(agentName string, cmd string) (string, error) {
	req := &dto.ExecuteRequest{
		AgentName: agentName,
		Command:   cmd,
	}

	var result string
	err := global.Client.Call("/handle/exec", req, &result, time.Second)
	return result, err
}
