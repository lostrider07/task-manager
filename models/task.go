package models

type Task struct {
	ID      string `json:"id" example:"1"`
	Title   string `json:"title" example:"Buy milk"`
	Details string `json:"details" example:"From the nearby store"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"not found"`
}