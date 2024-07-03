package hole

import (
	"github.com/sllt/hole/internal/agent/global"
	"github.com/sllt/nps/client"
	"github.com/sllt/nps/lib/config"
	"github.com/sllt/nps/lib/file"
	"sync"
)

func StartNpsClient(serverAddr string, vkey string) {
	cmConfig := &config.CommonConfig{
		Server:           serverAddr,
		VKey:             vkey,
		Tp:               "tcp",
		AutoReconnection: true,
		TlsEnable:        false,
		ProxyUrl:         "",
		Client: &file.Client{
			Cnf: &file.Config{
				U:        "",
				P:        "",
				Compress: false,
				Crypt:    false,
			},
			Id:              0,
			VerifyKey:       "",
			Addr:            "",
			Remark:          global.Agent.ShortName,
			Status:          false,
			IsConnect:       false,
			RateLimit:       1000,
			Flow:            nil,
			Rate:            nil,
			NoStore:         false,
			NoDisplay:       false,
			MaxConn:         0,
			NowConn:         0,
			WebUserName:     "",
			WebPassword:     "",
			ConfigConnAllow: false,
			MaxTunnelNum:    0,
			Version:         "",
			BlackIpList:     nil,
			LastOnlineTime:  "",
			RWMutex:         sync.RWMutex{},
		},
		DisconnectTime: 60,
	}

	conf := &config.Config{
		CommonConfig: cmConfig,
		Hosts:        nil,
		Tasks:        nil,
		Healths:      nil,
		LocalServer:  nil,
	}
	global.NpcStarted = true

	go client.StartFromConfig(conf)
}
