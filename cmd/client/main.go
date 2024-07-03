package main

import (
	_ "github.com/joho/godotenv/autoload"
	"os"

	"github.com/sllt/booby"
	blog "github.com/sllt/booby/log"
	"github.com/sllt/hole/cmd/client/cmd"
	"github.com/sllt/hole/internal/client/global"
	"github.com/sllt/log"
	"net"
	"time"
)

func main() {
	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = "localhost:3030"
	}
	log.SetReportCaller(true)
	blog.SetLevel(blog.LevelNone)
	client, err := booby.NewClient(func() (net.Conn, error) {
		return net.DialTimeout("tcp", serverAddr, time.Second*3)
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
