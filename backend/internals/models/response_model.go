package models

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   any `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}
