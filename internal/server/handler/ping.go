package handler

import (
	"github.com/sllt/booby"
)

func Ping(ctx *booby.Context) {
	_ = ctx.Write("ok")
}
