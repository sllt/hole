package dto

type ExitRequest struct {
	AgentName string `json:"agent_name"`
}

type ExecuteRequest struct {
	AgentName string `json:"agent_name"`
	Command   string `json:"command"`
}

type StartSocksRequest struct {
	Port      int    `json:"port"`
	AgentName string `json:"agent_name"`
}

type ShellRequest struct {
	AgentName string `json:"agent_name"`
}
