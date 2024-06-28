package dto

type ExitRequest struct {
	AgentName string `json:"agent_name"`
}

type ExecuteRequest struct {
	AgentName string `json:"agent_name"`
	Command   string `json:"command"`
}
