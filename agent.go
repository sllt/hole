package hole

import (
	"github.com/sllt/booby"
	"github.com/sllt/booby/log"
	"github.com/sllt/hole/internal/agent/global"
	"github.com/sllt/hole/internal/agent/router"
	"net"
	"time"
)

func StartAgent(serverAddr string) {

	log.SetLevel(log.LevelNone)
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

	client.Call("/agent/register", global.Agent, nil, time.Second)

	client.Set("short-name", global.Agent.ShortName)

	client.Handler.HandleConnected(func(client *booby.Client) {
		client.Call("/agent/register", global.Agent, nil, time.Second)
	})

}
