package global

import "hole/internal/model"

var (
	Agents map[string]*model.Agent
)

func init() {
	Agents = make(map[string]*model.Agent)
}
