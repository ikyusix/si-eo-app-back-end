package models

type Respons struct {
	Success bool
	Message string      `json:"successMessage"`
	Data    interface{} `json:"data,omitempty"`
}
