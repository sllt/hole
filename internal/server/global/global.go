package global

import "github.com/sllt/hole/internal/model"

var (
	Agents map[string]*model.Agent
)

func init() {
	Agents = make(map[string]*model.Agent)
}
