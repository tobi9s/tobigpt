package model

type ChatRequest struct {
	History string `json:"history"`
	Prompt  string `json:"prompt"`
}

type ChatResponse struct {
	History string `json:"history"`
	Prompt  string `json:"prompt"`
	Answer  string `json:"answer"`
}
