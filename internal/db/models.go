package db

import (
	"gorm.io/gorm"
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
	Session string `gorm:"primary_key;"`
	User_ID uint 
}
