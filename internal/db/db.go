package db

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	if err == nil {
		log.Println("Connected to Postgres Database")
	}
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

func (db *DB) SetSession(s Session) error {
	x := db.Create(&s)
	return x.Error
}

func (db *DB) DeleteSession(s Session) error {
	x := db.Where("session = ?", s.Session).Delete(&Session{})
	return x.Error
}
