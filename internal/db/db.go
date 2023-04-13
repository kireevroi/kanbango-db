package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Connect(URL string) error {
	gormdb, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	db.DB = gormdb
	return err
}

func (db *DB) CreateUser(u User) error {
	x := db.Create(&u)
	return x.Error
}

func (db *DB) GetUser(username string) (User, error) {
	var u User
	x := db.Where("username = ?", username).First(&u)
	return u, x.Error
}