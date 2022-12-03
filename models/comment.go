package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Comment string `json:"comment"`
}
