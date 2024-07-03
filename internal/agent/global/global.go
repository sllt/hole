package global

import (
	"github.com/gliderlabs/ssh"
	"github.com/sllt/go-socks5"
	"github.com/sllt/hole/internal/agent/service"
	"github.com/sllt/hole/internal/model"
)

var (
	Agent       *model.Agent
	NpcStarted  bool
	SocksError  string
	SocksServer *socks5.Server
	SshServer   *ssh.Server
)

func init() {
	agent, _ := service.GetAgentInfo()
	Agent = agent

	NpcStarted = false
}
