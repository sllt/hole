package handler

import (
	"github.com/sllt/booby"
	"os"
)

func Exit(ctx *booby.Context) {
	_ = ctx.Write("bye")
	os.Exit(0)
}

func Execute(ctx *booby.Context) {
	// TODO run cmd
	_ = ctx.Write("ok")
}
