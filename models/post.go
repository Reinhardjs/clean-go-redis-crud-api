package models

import (
	"time"
)

type Post struct {
	ID          int       `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (post *Post) Validate() (string, bool) {

	if post.Title == "" {
		return "Title should be on the payload", false
	}

	if post.Description == "" {
		return "Description should be on the payload", false
	}

	return "Payload is valid", true
}
