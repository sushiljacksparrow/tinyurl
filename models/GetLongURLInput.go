package models

type GetLongURLInput struct {
	TinyUrl		string `json:"tiny_url" binding:"required"`
	User 		string `binding:"required"`
}
