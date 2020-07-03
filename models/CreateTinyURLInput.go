package models

type CreateTinyURLInput struct {
	OriginalURL string `json:"original_url" binding:"required"`
	User 		string `json:"user" binding: "required"`
}
