package service

import (
	"github.com/sllt/af/random"
	"hole/internal/model"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetAgentInfo() (*model.Agent, error) {
	var sid string
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
		sid = random.RandString(5)
	} else {
		sid = strings.ToLower(hostname)
	}

	filename := filepath.Base(os.Args[0])

	sid = sid + "-" + filename + random.RandString(2)

	return &model.Agent{
		ShortName:   sid,
		Hostname:    hostname,
		OS:          runtime.GOOS,
		Description: "unknown",
	}, nil

}
