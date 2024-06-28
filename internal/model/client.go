package model

type Client struct {
	Name        string `json:"name"`
	ShortName   string `json:"short_name"`
	Description string `json:"description"`
}
