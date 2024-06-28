package main

import (
	"github.com/sllt/booby"
	blog "github.com/sllt/booby/log"
	"github.com/sllt/log"
	"hole/cmd/client/cmd"
	"hole/internal/client/global"
	"net"
	"time"
)

func main() {

	log.SetReportCaller(true)
	blog.SetLevel(blog.LevelNone)
	client, err := booby.NewClient(func() (net.Conn, error) {
		return net.DialTimeout("tcp", "localhost:3030", time.Second*3)
	})
	if err != nil {
		panic(err)
	}

	var result string
	client.Call("/ping", "hello", &result, time.Second)
	log.Info(result)
	if result != "ok" {
		log.Fatal("hole server error")
	}
	global.Client = client
	defer client.Stop()

	cmd.Execute()

}
