package models

import "time"

type URL struct {
	ID        	int64 	  `json:"id" gorm:"primary_key"`
	OriginalURL string 	  `json:"original_url"`
	TinyURL     string 	  `json:"tiny_url"`
	User 		string 	  `json:"user"`
	CreatedAt   time.Time `json:"created_at"`
}
