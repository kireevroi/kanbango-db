package db

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

// Migration models

type User struct {
	gorm.Model
	ID	uint `gorm:"primary_key"`
	Username string `gorm:"unique"`
	PasswordHash string	
}

type Session struct {
	gorm.Model
	Session uuid.UUID `gorm:"type:uuid;primary_key;"`
	User_ID uint 
}
