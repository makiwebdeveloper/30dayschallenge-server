package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;text;not null"`
	Password string `gorm:"text;not null"`
	Name     string `gorm:"text"`
}
