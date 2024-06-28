package model

import "github.com/sllt/booby"

type Agent struct {
	ShortName   string        `json:"short_name"`
	Hostname    string        `json:"hostname"`
	OS          string        `json:"os"`
	Description string        `json:"description"`
	Conn        *booby.Client `json:"-"`
}
