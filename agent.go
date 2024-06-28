package hole

import (
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/agent/router"
	"github.com/sllt/hole/internal/agent/service"
	"net"
	"time"
)

func StartAgent(serverAddr string) {
	client, err := booby.NewClient(func() (net.Conn, error) {
		return net.DialTimeout("tcp", serverAddr, time.Second*8)
	})
	//defer client.Stop()

	if err != nil {
		panic(err)
	}

	client.Handler.SetReconnectInterval(time.Second * 10)

	var result string
	client.Call("/ping", nil, &result, time.Second)
	router.RegisterAgentRoutes(client.Handler)

	agent, _ := service.GetAgentInfo()

	client.Call("/agent/register", &agent, nil, time.Second)

	client.Set("short-name", agent.ShortName)

	client.Handler.HandleConnected(func(client *booby.Client) {
		client.Call("/agent/register", &agent, nil, time.Second)
	})
}
