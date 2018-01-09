package models

import (
	"github.com/jinzhu/gorm"
)

// User struct for GORM model
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}
