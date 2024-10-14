package models

import "time"

type ApiResponse struct {
	Timestamp time.Time   `json:"timestamp"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func NewApiResponse() ApiResponse {
	return ApiResponse{
		Timestamp: time.Now(),
		Success:   true,
		Message:   "successful",
	}
}
