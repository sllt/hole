package handler

import (
	"fmt"
	"github.com/sllt/booby"
	"github.com/sllt/go-socks5"
	"github.com/sllt/hole/internal/agent/global"
	"github.com/sllt/hole/internal/model/dto"
	"log"
	"os"
)

func StartSocksServer(ctx *booby.Context) {
	var req dto.StartSocksRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.Write("param error")
		return
	}

	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
	)
	global.SocksServer = server

	go func() {
		if err := server.ListenAndServe("tcp", fmt.Sprintf(":%d", req.Port)); err != nil {
			log.Println(err)
			global.SocksError = err.Error()
		}
	}()

	ctx.Write("ok")

}

func StopSocksServer(ctx *booby.Context) {
	if global.SocksServer != nil {
		global.SocksServer.Stop()
		log.Println("stop socks server")
		global.SocksServer = nil
	}
	ctx.Write("ok")
}
